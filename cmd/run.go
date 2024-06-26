package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type Run struct {
}

var cmd *exec.Cmd

var lastEventMutex sync.Mutex

// Action
func (e Run) Action(c *cli.Context) error {
	main := c.Args().First()
	if len(main) == 0 {
		return errors.New("Please enter the executable file name")
	}
	// commandStr := "go build -o main.exe " + main + "&& main.exe"
	commandStr := "go run " + main
	err := runApp(commandStr)
	if err != nil {
		fmt.Println(err)
	}
	watchApp(commandStr)
	return nil
}

// watchApp
func watchApp(commandStr string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go watchForChanges(watcher)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			// 处理文件变化事件
			handleEvent(event, commandStr)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error:", err)
		}
	}
}

// handleEvent 处理文件系统事件
func handleEvent(event fsnotify.Event, commandStr string) {
	if event.Op&fsnotify.Write == fsnotify.Write {
		// 检查事件是否与上次相同
		lastEventMutex.Lock()
		defer lastEventMutex.Unlock()

		log.Println("File modified:", event.Name)

		// 先停止应用程序
		stopApp()

		// 重新启动应用程序
		runApp(commandStr)
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

	return nil
}

// stopApp 停止应用
func stopApp() {
	if runtime.GOOS == "windows" {
		var cmd2 *exec.Cmd
		cmd2 = exec.Command("cmd.exe", "/C", "TASKKILL /F /IM main.exe /T")
		if err := cmd2.Start(); err != nil {
			fmt.Println("Start error:", err)
		}
	} else if cmd != nil && cmd.Process != nil {
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("Error killing process:", err)
			return
		}
	}

	// 等待一段时间以确保进程已经结束
	time.Sleep(1 * time.Second)
}
