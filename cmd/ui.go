package cmd

import (
	"github.com/small-ek/ant-cli/boot/serve"
	"github.com/urfave/cli/v2"
)

type Ui struct {
}

// Action
func (e Ui) Action(c *cli.Context) error {
	serve.LoadSrv()
	return nil
}
