package command

import (
	"github.com/codegangsta/cli"
	"github.com/Atrox/homedir"
	"os"
	"io"
	"fmt"
)

func CmdInit(c *cli.Context) {
	path, err := homedir.Expand("~/.dailylog")
	if err != nil {
		panic(err.Error())
	}
	dest, err := os.Create(".dailylog")
	if err != nil {
		panic(err.Error())
	}
	defer dest.Close()
	if Exists(path) {
		src, err := os.Open(path)
		if err != nil {
			panic(err.Error())
		}
		defer src.Close()
		io.Copy(dest, src)
		fmt.Printf(".dailylog create from %s success\n", path)
	} else {
		io.Copy(dest, Initialfile)
		fmt.Println(".dailylog create success")
	}
	os.Exit(0)
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
