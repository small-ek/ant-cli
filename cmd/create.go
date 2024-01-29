package cmd

import (
	"github.com/urfave/cli/v2"
)

type Create struct {
}

func (e Create) Action(c *cli.Context) error {
	GenGo(c.Args().First())
	return nil
}
