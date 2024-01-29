package cmd

import (
	"fmt"
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
	// 确保父目录存在
	if err := os.MkdirAll(app, os.ModePerm); err != nil {
		fmt.Println("Error occurs when creating the root director:", err)
		return
	}

	//projectTree := TreePath{
	//	Name: app,
	//	Child: []TreePath{
	//		{Name: "main.go", Template: template.Main(app)},
	//		{
	//			Name: "app",
	//			Child: []TreePath{
	//				{Name: "dao"},
	//				{Name: "http"},
	//				{Name: "model"},
	//				{Name: "request"},
	//				{Name: "service"},
	//				{Name: "vo"},
	//			},
	//		},
	//		{
	//			Name: "boot",
	//			Child: []TreePath{
	//				{
	//					Name: "router",
	//					Child: []TreePath{
	//						{Name: "router.go", Template: template.Router(app)},
	//					},
	//				},
	//				{
	//					Name: "serve",
	//					Child: []TreePath{
	//						{Name: "serve.go", Template: template.Serve(app)},
	//					},
	//				},
	//			},
	//		},
	//		{
	//			Name: "config",
	//			Child: []TreePath{
	//				{
	//					Name:     "config.toml",
	//					Template: template.Config(app),
	//				},
	//			},
	//		},
	//		{
	//			Name: "log",
	//			Child: []TreePath{
	//				{
	//					Name: "ant.log",
	//				},
	//			},
	//		},
	//		{
	//			Name: "router",
	//			Child: []TreePath{
	//				{
	//					Name:     "index.go",
	//					Template: template.RouterIndex(app),
	//				},
	//			},
	//		},
	//	},
	//}
	projectTree := TreePath{
		Name: app,
		Child: []TreePath{
			{
				Name:     "cmd",
				Template: "package main\n\nfunc main() {\n\t// 你的应用逻辑在这里\n}\n",
			},
			{
				Name: "pkg",
				Child: []TreePath{
					{
						Name:     "util",
						Template: "package util\n\n// 你的工具函数在这里\n",
					},
				},
			},
		},
	}

	err := generateFiles(app, projectTree)
	if err != nil {
		fmt.Println("错误:", err)
		return
	}

	fmt.Println("cd " + app)
	fmt.Println("go mod tidy")
	fmt.Println("go mod vendor")
}

// generateFiles 生成代码
func generateFiles(rootPath string, tree TreePath) error {
	currentPath := filepath.Join(rootPath, tree.Name)
	fmt.Println(currentPath)
	fmt.Println(tree.Name)
	if err := os.Mkdir(currentPath, os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}

	for _, child := range tree.Child {
		if err := generateFiles(currentPath, child); err != nil {
			return err
		}
	}

	if tree.Template != "" {
		fileName := strings.ToLower(tree.Name) + ".go" // Assuming .go files for simplicity
		filePath := filepath.Join(currentPath, fileName)
		if _, err := os.Create(filePath); err != nil {
			return err
		}

		err := os.WriteFile(filePath, []byte(tree.Template), os.ModePerm)
		if err != nil {
			return err
		}
		//fmt.Printf("File created: %s\n", filePath)
	}

	return nil
}
