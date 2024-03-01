package cmd

import (
	"errors"
	"github.com/urfave/cli/v2"
)

type GenApi struct {
}

// Action
func (b GenApi) Action(c *cli.Context) error {
	main := c.Args().First()
	if len(main) == 0 {
		return errors.New("Please enter the database name and table name")
	}
	
	return nil
}
