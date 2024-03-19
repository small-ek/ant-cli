package main

import (
	"embed"
	"flag"
	"github.com/small-ek/ant-cli/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

type Func struct {
	Env      cmd.Env
	Create   cmd.Create
	Run      cmd.Run
	Build    cmd.Build
	Install  cmd.Install
	GenDao   cmd.GenDao
	GenApi   cmd.GenApi
	GenModel cmd.GenModel
	Ui       cmd.Ui
}

// 嵌入前端构建的静态资源
//
//go:embed dist/*
var frontendEmbed embed.FS

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
				Action:  funcs.Install.Action,
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
				Action:  funcs.Run.Action,
			},
			{
				Name:    "ui",
				Aliases: []string{"u"},
				Usage:   "Run go code with hot compilation-like features",
				Action:  funcs.Ui.Action,
			},
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build Go projects cross-platform",
				Action:  funcs.Build.Action,
			},
			{
				Name:    "gen",
				Aliases: []string{"g"},
				Usage: `The "GEN" command is used for multiple generation purposes.
It is based on database generation controller, model, service, router, API`,
				Subcommands: []*cli.Command{
					{
						Name:   "model",
						Usage:  "Automatically generate Model files for SQL",
						Action: funcs.GenModel.Action,
					},
					{
						Name:   "dao",
						Usage:  "Automatically generate DAO and Model files for SQL",
						Action: funcs.GenDao.Action,
					},
					{
						Name:   "api",
						Usage:  "Automatically generate API interface files for SQL",
						Action: funcs.GenApi.Action,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
