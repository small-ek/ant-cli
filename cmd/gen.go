package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

type Gen struct {
}

// Action
func (b Gen) Action(c *cli.Context) error {
	main := c.Args().First()
	if len(main) == 0 {
		return errors.New("Please enter the database name and table name")
	}
	// 构建命令字符串
	commandStr := "go"
	args := []string{"build", "-o", main, "-ldflags", "-s -w"}
	err := buildApp(commandStr, args)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
