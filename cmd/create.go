package cmd

import (
	"fmt"
	"github.com/small-ek/ant-cli/aexec"
	"github.com/urfave/cli/v2"
	"log"
	"strings"
)

type Create struct {
}

func (e Create) Action(c *cli.Context) error {
	result, err := aexec.ShellExec("go env")
	if err != nil {
		log.Println(err)
	}
	if result == "" {
		log.Println(`Failed to get Golang environment variables`)
	}
	env := strings.Split(result, "\n")
	for _, row := range env {
		row = strings.Trim(row, " ")
		if row == "" {
			continue
		}
		fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println(row)
	}
	return nil
}
