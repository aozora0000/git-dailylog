package command

import (
	"github.com/codegangsta/cli"
	"io/ioutil"
		"os/exec"
	"fmt"
	"os"
	"strings"
	)

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
		for i, _ := range s.lines {
			ret <- s.lines[len(s.lines)-1-i]
		}
		close(ret)
	}()
	return ret
}

// git log --after="2015-09-25 00:00:00" --before="2015-09-26 00:00:00" --date=local --pretty=format:"%h: %ad %an: %s" --author "Kazuhiko Hotta"
func CmdGet(c *cli.Context) {
	config, err := ioutil.ReadFile(".dailylog")
	if err != nil {
		panic(err)
	}

	parser := &TimeDurationParser{c.String("ago")}
	author := c.String("author")

	timestamps := parser.Parse()
	var args = []string{
		"log",
		"--date=iso",
		"--pretty=tformat:" + string(config),
		"--after=\"" + timestamps.From.String() + "\"",
		"--before==\"" + timestamps.To.String() + "\"",
	}
	if author != "" {
		args = append(args, "--author=\"" + author + "\"")
	}

	out, _ := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	lines := &Lines{strings.Split(string(out), "\n")}
	for line := range lines.Get(c.Bool("reverse")) {
		fmt.Println(strings.Replace(strings.Replace(strings.Trim(line, "\""), "\\n", ":break:", -1), ":break:", `
`, -1))
	}
	os.Exit(0)
}
