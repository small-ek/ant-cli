package main

import (
	"flag"
	"fmt"
	"github.com/small-ek/ant-cli/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type Func struct {
	Env    cmd.Env
	Create cmd.Create
}

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	flag.Parse()
	funcs := Func{}
	app := &cli.App{
		Name:    "ant-cli",
		Usage:   "Used to build antgo projects",
		Version: "1.0.0",
		Flags:   []cli.Flag{},
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

					commandStr := "go mod tidy"
					var cmd *exec.Cmd

					if runtime.GOOS == "windows" {
						cmd = exec.Command("cmd.exe", "/C", commandStr)
					} else {
						cmd = exec.Command("sh", "-c", commandStr)
					}

					// Run the command and get the output
					output, err := cmd.CombinedOutput()
					if err != nil {
						panic(err)
					}

					// Print the output
					fmt.Println(string(output))

					return nil
				},
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create an Antgo application",
				Action:  funcs.Create.Action,
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
				Aliases: []string{"b"},
				Usage:   "Build Go projects cross-platform",
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args().First())
					return nil
				},
			},
			{
				Name:    "docker",
				Aliases: []string{"d"},
				Usage:   "Build a docker image for the current Antgo project",
				Action: func(c *cli.Context) error {
					fmt.Println(c.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
