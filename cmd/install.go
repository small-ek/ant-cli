package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
	"runtime"
)

type Install struct {
}

func (e Install) Action(c *cli.Context) error {
	var cmd1 *exec.Cmd
	var cmd2 *exec.Cmd
	tidy := "go mod tidy"

	if runtime.GOOS == "windows" {
		cmd1 = exec.Command("cmd.exe", "/C", tidy)
	} else {
		cmd1 = exec.Command("sh", "-c", tidy)
	}
	_, err := cmd1.CombinedOutput()
	if err != nil {
		panic(err)
	}

	vendor := "go mod vendor"
	if runtime.GOOS == "windows" {
		cmd2 = exec.Command("cmd.exe", "/C", vendor)

	} else {
		cmd2 = exec.Command("sh", "-c", vendor)
	}
	_, err2 := cmd2.CombinedOutput()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("Successful installation")

	return nil
}
