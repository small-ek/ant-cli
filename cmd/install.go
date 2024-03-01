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
	commandStr := "go mod tidy && go mod vendor"
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", commandStr)
	} else {
		cmd = exec.Command("sh", "-c", commandStr)
	}

	_, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful installation")

	return nil
}
