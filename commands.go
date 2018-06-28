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
		Usage:  "Initialize dailylog fotmat file. format reference: https://git-scm.com/docs/pretty-formats",
		Action: command.CmdInit,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "get",
		Usage:  "Get abarge commit log formatted from .dailylog",
		Action: command.CmdGet,
		Flags:  []cli.Flag{
			cli.StringFlag{
				Name: "author",
				Value: "",
				Usage: "commit author filter",
			},
			cli.BoolFlag{
				Name: "reverse",
				Usage: "reverse output",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
