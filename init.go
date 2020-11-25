package git_dailylog

import (
	"fmt"
	"github.com/Atrox/homedir"
	"github.com/k0kubun/pp"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"path/filepath"
)

var CmdInitCommand = &cli.Command{
	Name:   "init",
	Usage:  "Initialize dailylog fotmat file. format reference: https://git-scm.com/docs/pretty-formats",
	Action: CmdInit,
	Flags:  []cli.Flag{},
}

func CmdInit(c *cli.Context) error {
	path, err := homedir.Expand("~/.dailylog")
	if err != nil {
		pp.Println(err.Error())
		return err
	}
	rootPath, err := getRoot()
	if err != nil {
		pp.Println(err.Error())
		return err
	}
	dest, err := os.Create(filepath.Join(rootPath, ".dailylog"))
	if err != nil {
		pp.Println(err.Error())
		return err
	}
	defer dest.Close()
	if Exists(path) {
		src, err := os.Open(path)
		if err != nil {
			pp.Println(err.Error())
			return err
		}
		defer src.Close()
		io.Copy(dest, src)
		fmt.Printf("%s create from %s success\n", filepath.Join(rootPath, ".dailylog"), path)
	} else {
		io.Copy(dest, Initialfile)
		fmt.Println(".dailylog create success")
	}
	return nil
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
