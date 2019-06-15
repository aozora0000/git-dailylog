package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/k0kubun/pp"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var CmdGetCommand = cli.Command{
	Name:   "get",
	Usage:  "Get abarge commit log formatted from .dailylog",
	Action: CmdGet,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "author",
			Value: "",
			Usage: "commit author filter",
		},
		cli.BoolFlag{
			Name:  "reverse",
			Usage: "reverse output",
		},
	},
}

type Lines struct {
	lines []string
}

func (s *Lines) Get(isReverse bool) chan string {
	if isReverse {
		return s.Reverse()
	} else {
		return s.Normal()
	}
}

func (s *Lines) Normal() chan string {
	ret := make(chan string)
	go func() {
		for _, line := range s.lines {
			ret <- line
		}
		close(ret)
	}()
	return ret
}

func (s *Lines) Reverse() chan string {
	ret := make(chan string)
	go func() {
		for i := range s.lines {
			ret <- s.lines[len(s.lines)-1-i]
		}
		close(ret)
	}()
	return ret
}

// git log --after="2015-09-25 00:00:00" --before="2015-09-26 00:00:00" --date=local --pretty=format:"%h: %ad %an: %s" --author "Kazuhiko Hotta"
func CmdGet(c *cli.Context) {
	rootPath, err := getRoot()
	if err != nil {
		pp.Println(err.Error())
		os.Exit(1)
	}
	config, err := ioutil.ReadFile(filepath.Join(rootPath, ".dailylog"))
	if err != nil {
		pp.Println(err.Error())
		os.Exit(1)
	}
	var ago = "today"

	if c.Args().Get(0) != "" {
		ago = c.Args().Get(0)
	}

	parser := &TimeDurationParser{ago}
	author := c.String("author")

	timestamps := parser.Parse()

	var args = []string{
		"log",
		"--date=iso",
		"--branches",
		"--tags",
		"--pretty=format:" + string(config),
		"--after=\"" + timestamps.From.String() + "\"",
		"--before=\"" + timestamps.To.String() + "\"",
	}
	if author != "" {
		args = append(args, "--author="+author)
	}

	out, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		pp.Println(err.Error())
		os.Exit(1)
	}
	lines := &Lines{strings.Split(string(out), "\n")}
	for line := range lines.Get(c.Bool("reverse")) {
		fmt.Println(strings.Trim(line, "\""))
	}
	os.Exit(0)
}
