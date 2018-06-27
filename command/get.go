package command

import (
	"github.com/codegangsta/cli"
		//"os/exec"
	"io/ioutil"
	"github.com/k0kubun/pp"
)

// git log --after="2015-09-25 00:00:00" --before="2015-09-26 00:00:00" --date=local --pretty=format:"%h: %ad %an: %s" --author "Kazuhiko Hotta"
func CmdGet(c *cli.Context) {
	config, err := ioutil.ReadFile(".dailylog")
	if err != nil {
		panic(err)
	}

	parser := &TimeDurationParser{c.String("day")}
	author := c.String("author")

	timestamps := parser.Parse()
	var args = []string{
		"log",
		"--date", "local",
		"--pretty", "format:" + string(config),
		"--before", timestamps.From.String(),
		"--after", timestamps.To.String(),
	}
	if author != "" {
		args = append(args, "--author")
		args = append(args, author)
	}
	pp.Println(args)
	//exec.Command("git", args)

}
