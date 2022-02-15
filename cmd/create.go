package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

type Create struct {
}

func (e Create) Action(c *cli.Context) error {
	fmt.Println(c.Args().First())

	return nil
}
