package command

import (
	"github.com/codegangsta/cli"
	"os"
	"io"
	"fmt"
)

func CmdInit(c *cli.Context) {
	file, err := os.Create(".dailylog")
	if err != nil {
		panic(err.Error())
	}
	io.Copy(file, Initialfile)
	defer file.Close()

	fmt.Println(".dailylog create success")
	os.Exit(0)
}
