package cmd

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

type Run struct {
}

var cmd *exec.Cmd
var cmdString string // 保存当前正在运行的命令字符串

func (e Run) Action(c *cli.Context) error {
	main := c.Args().First()
	commandStr := "go run " + main
	err := runApp(commandStr)
	if err != nil {
		fmt.Println(err)
	}
	watchApp(commandStr)
	return nil
}

func watchApp(commandStr string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go watchForChanges(watcher)

	// 主循环
	for {
		select {
		case event := <-watcher.Events:
			// 处理文件变化事件
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("File modified:", event.Name)
				time.Sleep(5 * time.Second)

				// 先停止监听端口
				stopApp()

				// 重新启动应用
				runApp(commandStr)
			}
		case err := <-watcher.Errors:
			log.Println("Error:", err)
		}
	}
}

func watchForChanges(watcher *fsnotify.Watcher) {
	// 监听当前目录下的所有Go文件
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".go" {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

// runApp 启动应用
func runApp(commandStr string) error {
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", commandStr)
	} else {
		cmd = exec.Command("sh", "-c", commandStr)
	}

	// 创建管道获取标准输出和标准错误输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 启动命令
	err = cmd.Start()
	if err != nil {
		fmt.Println("Start error:", err)
		return err
	}

	// 实时读取标准输出
	go func() {
		reader := bufio.NewReader(stdout)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return
			}
			fmt.Print(line)
		}
	}()

	// 实时读取标准错误输出
	go func() {
		reader := bufio.NewReader(stderr)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return
			}
			fmt.Print(line)
		}
	}()

	// 保存当前命令字符串
	cmdString = commandStr

	return nil
}

// stopApp 停止应用
func stopApp() {
	if cmd != nil && cmd.Process != nil {
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("Error killing process:", err)
			return
		}
		// 等待一段时间以确保进程已经结束
		time.Sleep(2 * time.Second)
	}
	// 清空当前命令字符串
	cmdString = ""
}
