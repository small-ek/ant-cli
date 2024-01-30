package cmd

import (
	"fmt"
	"github.com/small-ek/ant-cli/template"
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
	//if err := os.MkdirAll(app, os.ModePerm); err != nil {
	//	fmt.Println("Error occurs when creating the root director:", err)
	//	return
	//}

	projectTree := TreePath{
		Name: app,
		Child: []TreePath{
			{Name: "main.go", Template: template.Main(app)},
			{
				Name: "app",
				Child: []TreePath{
					{Name: "dao"},
					{Name: "http"},
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
				Name: "log",
				Child: []TreePath{
					{
						Name: "ant.log",
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

	err := generateFiles("./", projectTree)
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
	fmt.Println("--------------")
	fmt.Println(rootPath)
	fmt.Println(currentPath)
	// 创建当前目录
	if err := os.Mkdir(currentPath, os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}
	// 写入文件
	if tree.Template != "" {
		fileName := strings.ToLower(tree.Name) // 示例中文件名为小写的目录名
		filePath := filepath.Join(currentPath, fileName)
		fmt.Println(filePath)
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		// 写入模板内容
		_, err = file.WriteString(tree.Template)
		if err != nil {
			return err
		}
		fmt.Printf("文件 %s 已成功创建。\n", filePath)
	}
	// 递归创建子目录
	for _, child := range tree.Child {
		if err := generateFiles(currentPath, child); err != nil {
			return err
		}
	}

	return nil
}
