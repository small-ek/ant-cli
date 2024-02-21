package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

type Build struct {
}

// Action
func (b Build) Action(c *cli.Context) error {
	main := c.Args().First()
	if len(main) == 0 {
		return errors.New("Please enter the compiled file name")
	}
	// 构建命令字符串
	commandStr := "go"
	args := []string{"build", "-o", main, "-ldflags", "-s -w"}
	err := buildApp(commandStr, args)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// buildApp 打包应用
func buildApp(command string, args []string) error {
	if runtime.GOOS == "windows" {
		cmd = exec.Command(command, args...)
	} else {
		cmd = exec.Command("sh", "-c", command+" "+quoteArgs(args))
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(output)) // 输出错误信息
		return err
	}
	return nil
}

// quoteArgs 将参数用引号括起来
func quoteArgs(args []string) string {
	for i, arg := range args {
		args[i] = fmt.Sprintf(`"%s"`, arg)
	}
	return strings.Join(args, " ")
}
