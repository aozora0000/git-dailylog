package main

import (
	"fmt"
	"github.com/aozora0000/git-dailylog/command"
	"os"

	"github.com/codegangsta/cli"
)

var Name = "git-dailylog"
var Version = "0.4.0"
var GlobalFlags = []cli.Flag{}

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "aozora0000"
	app.Email = "aozora0000@gmail.com"
	app.Usage = "Garbage Commit Log."

	app.Flags = GlobalFlags
	app.Commands = []cli.Command{
		command.CmdGetCommand,
		command.CmdInitCommand,
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
		os.Exit(2)
	}

	app.Run(os.Args)
}
