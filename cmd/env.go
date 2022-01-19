package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

type Env struct {
}

func (e Env) Action(c *cli.Context) error {
	fmt.Println(c.Args().First())
	return nil
}
