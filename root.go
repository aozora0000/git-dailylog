package git_dailylog

import (
	"os/exec"
	"strings"
)

func getRoot() (string, error) {
	var args = []string{
		"rev-parse",
		"--show-toplevel",
	}
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(out), "\n"), nil
}
