package main

import (
	"fmt"
	"os"

	"github.com/aozora0000/git-dailylog/command"
	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "init",
		Usage:  "Initialize Dailylog Fotmat File",
		Action: command.CmdInit,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "get",
		Usage:  "",
		Action: command.CmdGet,
		Flags:  []cli.Flag{
			cli.StringFlag{
				Name: "ago",
				Value: "today",
				Usage: "gabarge commit log at day. [today, 7day, 1week, 1years]",
			},
			cli.StringFlag{
				Name: "author",
				Value: "",
				Usage: "commit author filter",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}