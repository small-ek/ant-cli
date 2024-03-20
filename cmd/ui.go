package cmd

import (
	"embed"
	"github.com/small-ek/ant-cli/boot/serve"
	"github.com/urfave/cli/v2"
)

type Ui struct {
	Fs embed.FS
}

// Action
func (u Ui) Action(c *cli.Context) error {
	serve.LoadSrv(u.Fs)
	return nil
}
