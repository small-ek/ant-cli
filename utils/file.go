package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

// Exists 判断路径是否存在文件
func Exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// GetFileName
func GetFileName() string {
	// 获取当前工作目录的路径
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败:", err)
	}

	// 使用 filepath.Base 获取路径的基础名称
	dirName := filepath.Base(currentDir)

	return dirName
}
