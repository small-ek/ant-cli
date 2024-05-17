package cmd

import (
	"fmt"
	"github.com/small-ek/ant-cli/template"
	"github.com/small-ek/ant-cli/utils"
	"os"
	"path/filepath"
	"strings"
)

type TreePath struct {
	Name     string
	Child    []TreePath
	Template string
}

func GenGo(app string) {
	projectTree := TreePath{
		Name: app,
		Child: []TreePath{
			{Name: "main.go", Template: template.Main(app)},
			{Name: "go.mod", Template: template.Mod(app)},
			{Name: "README.md", Template: template.Readme()},
			{
				Name: "app",
				Child: []TreePath{
					{Name: "dao"},
					{Name: "http", Child: []TreePath{
						{Name: "base.go", Template: template.Base(app)},
						{Name: "index", Child: []TreePath{
							{Name: "index.go", Template: template.Index(app)},
						}},
					}},
					{Name: "model"},
					{Name: "request"},
					{Name: "service"},
					{Name: "vo"},
				},
			},
			{
				Name: "boot",
				Child: []TreePath{
					{
						Name: "router",
						Child: []TreePath{
							{Name: "router.go", Template: template.Router(app)},
						},
					},
					{
						Name: "serve",
						Child: []TreePath{
							{Name: "serve.go", Template: template.Serve(app)},
						},
					},
				},
			},
			{
				Name: "config",
				Child: []TreePath{
					{
						Name:     "config.toml",
						Template: template.Config(app),
					},
				},
			},
			{
				Name: "docs",
				Child: []TreePath{
					{
						Name:     "docs.go",
						Template: template.Docs(app),
					},
				},
			},
			{
				Name: "log",
				Child: []TreePath{
					{
						Name:     "ant.log",
						Template: "",
					},
				},
			},
			{
				Name: "router",
				Child: []TreePath{
					{
						Name:     "index.go",
						Template: template.RouterIndex(app),
					},
				},
			},
		},
	}

	createProjectTree(projectTree, ".")
	fmt.Println("Successful creation " + app)
	fmt.Println("cd " + app)
	fmt.Println("ant-cli install")
	fmt.Println("ant-cli run main.go")
}

func createProjectTree(node TreePath, parentPath string) {
	currentPath := filepath.Join(parentPath, node.Name)

	err := os.MkdirAll(currentPath, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory %s: %v\n", currentPath, err)
		return
	}

	for _, child := range node.Child {
		if child.Template != "" || strings.Contains(child.Name, ".") {
			filePath := filepath.Join(currentPath, child.Name)
			err := utils.WriteFile(filePath, child.Template)
			if err != nil {
				fmt.Printf("Error creating file %s: %v\n", filePath, err)
			}
		} else {
			createProjectTree(child, currentPath)
		}
	}
}
