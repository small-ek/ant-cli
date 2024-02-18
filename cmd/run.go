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

func (e Run) Action(c *cli.Context) error {
	main := c.Args().First()
	commandStr := "go run " + main
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", commandStr)
	} else {
		cmd = exec.Command("sh", "-c", commandStr)
	}

	// 创建管道获取标准输出和标准错误输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("无法获取标准输出管道:", err)
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("无法获取标准错误输出管道:", err)
		return err
	}

	// 启动命令
	err = cmd.Start()
	if err != nil {
		fmt.Println("命令启动失败:", err)
		return err
	}

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
				restartApp()
			}
		case err := <-watcher.Errors:
			log.Println("Error:", err)
		}
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

	// 等待命令执行完成
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
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

func restartApp() {
	log.Println("Restarting application...")
	cmd := exec.Command("go", "run", ".") // 这里使用了简单的重新运行方式，实际生产中可以替换为编译命令并重启
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Println("Error restarting application:", err)
	}
	// 在实际生产中，你可能需要等待一段时间，确保新的进程完全启动
	time.Sleep(2 * time.Second)
	os.Exit(0) // 退出当前进程，让新的进程接管
}
