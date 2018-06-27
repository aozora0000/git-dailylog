package command

import (
	"github.com/codegangsta/cli"
	"io/ioutil"
		"os/exec"
	"fmt"
	"os"
	"strings"
)

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
		"--pretty=format:" + string(config),
		"--after=\"" + timestamps.From.String() + "\"",
		"--before==\"" + timestamps.To.String() + "\"",
	}
	if author != "" {
		args = append(args, "--author=\"" + author + "\"")
	}
	fmt.Println(strings.Join(args, " "))
	out, _ := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		fmt.Println(strings.Trim(line, "\""))
	}
	os.Exit(0)
}
