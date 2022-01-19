package main

import (
	"fmt"
	"github.com/small-ek/ant-cli/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type Func struct {
	Env cmd.Env
}

func main() {
	funcs := Func{}
	app := &cli.App{
		Name:    "ant-cli",
		Usage:   "Used to build antgo projects",
		Version: "1.0.0",
		Flags: []cli.Flag{

		},
		Action: func(c *cli.Context) error {
			//fmt.Println(c.String("lang"))
			//fmt.Println(c.String("aaa"))
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Install ant binary to system environment variables (requires run permission)",
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args().First())
					return nil
				},
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create an Antgo application",
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args().First())
					return nil
				},
			},
			{
				Name:    "env",
				Aliases: []string{"e"},
				Usage:   "Show current Golang environment variables",
				Action:  funcs.Env.Action,
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Run go code with hot compilation-like features",
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args().First())
					return nil
				},
			},
			{
				Name:    "build",
				Aliases: []string{"r"},
				Usage:   "Build Go projects cross-platform",
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args().First())
					return nil
				},
			},
			{
				Name:    "docker",
				Aliases: []string{"r"},
				Usage:   "Build a docker image for the current Antgo project",
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args().First())
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
