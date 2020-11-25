package main

import (
	"fmt"
	git_dailylog "github.com/aozora0000/git-dailylog"
	"github.com/blang/semver"
	"github.com/codegangsta/cli"
	"github.com/pkg/errors"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"log"
	"os"
)

var Name = "git-dailylog"
var Version = "0.5.0"
var GlobalFlags = []cli.Flag{}
var Slug = "aozora0000/git-dailylog"

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "aozora0000"
	app.Email = "aozora0000@gmail.com"
	app.Usage = "Garbage Commit Log."

	app.Flags = GlobalFlags
	app.Commands = []cli.Command{
		git_dailylog.CmdGetCommand,
		git_dailylog.CmdInitCommand,
		selfUpdateCommand,
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
		os.Exit(2)
	}

	app.Run(os.Args)
}

func DoSelfUpdate(version string, slug string) (bool, error) {
	v := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(v, slug)
	return !latest.Version.Equals(v), errors.Wrap(err, "Binary update failed")
}

var selfUpdateCommand = cli.Command{
	Name:  "selfupdate",
	Usage: "Get abarge commit log formatted from .dailylog",
	Action: func(context *cli.Context) error {
		updated, err := DoSelfUpdate(Version, Slug)
		if err != nil {
			return err
		}
		if updated {
			log.Println("Current binary is the latest version", Version)
		} else {
			log.Println("Successfully updated to version", Version)
		}
		return nil
	},
	Flags: []cli.Flag{},
}
