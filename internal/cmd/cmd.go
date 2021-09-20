package cmd

import (
	"os"
	"os/exec"
	"strings"
)

func ExecCmd(dir string, cmd string) error {
	cmds := strings.Split(cmd, " ")

	c := exec.Command(cmds[0], cmds[1:]...)

	if dir != "" {
		c.Dir = dir
	}

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}

func ExecCmdOutput(dir string, cmd string) ([]byte, error) {
	cmds := strings.Split(cmd, " ")

	c := exec.Command(cmds[0], cmds[1:]...)

	if dir != "" {
		c.Dir = dir
	}

	return c.Output()
}
