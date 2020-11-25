package main

import (
	"fmt"
	git_dailylog "github.com/aozora0000/git-dailylog"
	"github.com/urfave/cli/v2"
	"os"
)

var Name = "git-dailylog"
var Version = "unknown"

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Authors = []*cli.Author{&cli.Author{Name: "aozora0000", Email: "aozora0000@gmail.com"}}
	app.Usage = "Garbage Commit Log."
	app.Metadata = map[string]interface{}{"Slug": "aozora0000/git-dailylog"}
	app.Flags = []cli.Flag{}
	app.Commands = []*cli.Command{
		git_dailylog.CmdGetCommand,
		git_dailylog.CmdInitCommand,
		git_dailylog.SelfUpdateCommand,
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
		os.Exit(2)
	}

	app.Run(os.Args)
}
