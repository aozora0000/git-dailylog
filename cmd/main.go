package main

import (
	"bufio"
	"errors"
	"fmt"
	git_dailylog "github.com/aozora0000/git-dailylog"
	"github.com/blang/semver"
	"github.com/codegangsta/cli"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"log"
	"os"
)

var Name = "git-dailylog"
var Version = "0.6.0"
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

var selfUpdateCommand = cli.Command{
	Name:  "selfupdate",
	Usage: "latest update from server",
	Action: func(context *cli.Context) error {
		if context.Bool("verbose") {
			selfupdate.EnableLog()
		}

		latest, found, err := selfupdate.DetectLatest(Slug)
		if err != nil {
			log.Println("Error occurred while detecting version:", err)
			return err
		}
		v := semver.MustParse(Version)
		if !found || latest.Version.LTE(v) {
			log.Println("Current version is the latest")
			return nil
		}

		fmt.Print("Do you want to update to ", latest.Version, "? (y/n): ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		if input != "y\n" && input != "n\n" {
			return errors.New("Invalid input")
		}
		if input == "n\n" {
			return nil
		}

		exe, err := os.Executable()
		if err != nil {
			return err
		}
		if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
			return err
		}
		log.Println("Successfully updated to version", latest.Version)
		return nil
	},
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose,v",
			Usage: "Verbose",
		},
	},
}
