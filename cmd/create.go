package cmd

import (
	"errors"
	"github.com/urfave/cli/v2"
)

type Create struct {
}

func (e Create) Action(c *cli.Context) error {
	name := c.Args().First()
	if len(name) == 0 {
		return errors.New("Please enter the project name")
	}
	GenGo(name)
	return nil
}
