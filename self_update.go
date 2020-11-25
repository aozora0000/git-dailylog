package git_dailylog

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var SelfUpdateCommand = &cli.Command{
	Name:  "selfupdate,self-update",
	Usage: "latest update from server",
	Action: func(context *cli.Context) error {
		if context.Bool("verbose") {
			selfupdate.EnableLog()
		}
		latest, found, err := selfupdate.DetectLatest(context.App.Metadata["Slug"].(string))
		if err != nil {
			log.Println("Error occurred while detecting version:", err)
			return err
		}
		v := semver.MustParse(context.App.Version)
		if !found || latest.Version.LTE(v) {
			log.Println("Current version is the latest")
			return nil
		}

		fmt.Print("Do you want to update to ", latest.Version, " ? (y/n): ")
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
		&cli.BoolFlag{
			Name:  "verbose,v",
			Usage: "Verbose",
		},
	},
}
