package cmd

import (
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os/exec"
)

type Env struct {
}

func (e Env) Action(c *cli.Context) error {
	cmd := exec.Command("cmd")
	in := bytes.NewBuffer(nil)

	cmd.Stdin = in
	var out bytes.Buffer
	cmd.Stdout = &out
	data, _ := cmd.Output()
	log.Println(data)
	go func() {
		// start stop restart
		in.WriteString("go env\n") //写入你的命令，可以有多行，"\n"表示回车
	}()
	if err := cmd.Start(); err != nil {
		log.Println(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Println(err)
	}
	rt := out.String()
	fmt.Println(rt)
	return nil
}
