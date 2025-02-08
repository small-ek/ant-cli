package main

import (
	"embed"
	_ "embed"
	"flag"
	"github.com/small-ek/ant-cli/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// 嵌入前端构建的静态资源
// Embed front-end built static resources
//
//go:embed web/dist/*
var frontendAssets embed.FS

// Func 结构体包含所有命令的处理函数
// Func struct contains all command handlers
type CommandHandlers struct {
	Env      cmd.Env
	Create   cmd.Create
	Run      cmd.Run
	Build    cmd.Build
	Install  cmd.Install
	GenDao   cmd.GenDao
	GenApi   cmd.GenApi
	GenModel cmd.GenModel
	Ui       cmd.Ui
	Rsa      cmd.Rsa
}

func main() {
	// 设置日志输出格式
	// Set log output format
	log.SetFlags(log.Llongfile | log.LstdFlags)

	// 解析命令行参数
	// Parse command line arguments
	flag.Parse()

	// 初始化命令处理函数
	// Initialize command handlers
	handlers := CommandHandlers{
		Ui: cmd.Ui{Fs: frontendAssets},
	}

	// 创建 CLI 应用
	// Create CLI application
	app := &cli.App{
		Name:    "ant-cli",
		Usage:   "Used to build antgo projects",
		Version: "1.0.3",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "Install ant binary to system environment variables (requires run permission)",
				Action:  handlers.Install.Action,
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create an Antgo application",
				Action:  handlers.Create.Action,
			},
			{
				Name:    "env",
				Aliases: []string{"e"},
				Usage:   "Show current Golang environment variables",
				Action:  handlers.Env.Action,
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Run go code with hot compilation-like features",
				Action:  handlers.Run.Action,
			},
			{
				Name:    "ui",
				Aliases: []string{"u"},
				Usage:   "Run go code with hot compilation-like features",
				Action:  handlers.Ui.Action,
			},
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build Go projects cross-platform",
				Action:  handlers.Build.Action,
			},
			{
				Name:  "rsa",
				Usage: "Generate RSA certificate",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "size",
						Usage: "Size of the RSA key",
						Value: 2048, // Default value
					},
				},
				Action: handlers.Rsa.Action,
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
						Action: handlers.GenModel.Action,
					},
					{
						Name:   "dao",
						Usage:  "Automatically generate DAO and Model files for SQL",
						Action: handlers.GenDao.Action,
					},
					{
						Name:   "api",
						Usage:  "Automatically generate API interface files for SQL",
						Action: handlers.GenApi.Action,
					},
				},
			},
		},
	}

	// 运行 CLI 应用
	// Run CLI application
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
