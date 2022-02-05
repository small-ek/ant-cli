package cmd

import (
	"github.com/small-ek/ant-cli/aexec"
	"github.com/urfave/cli/v2"
	"log"
	"strings"
)

type Env struct {
}

func (e Env) Action(c *cli.Context) error {
	result, err := aexec.ShellExec("go env")
	if err != nil {
		log.Println(err)
	}
	if result == "" {
		log.Println(`Failed to get Golang environment variables`)
	}
	env := strings.Split(result, "\n")
	for _, row := range env {
		row = strings.Trim(row, " ")
		if row == "" {
			continue
		}
		if gstr.Pos(line, "set ") == 0 {
			line = line[4:]
		}
		match, _ := gregex.MatchString(`(.+?)=(.*)`, line)
		if len(match) < 3 {
			mlog.Fatalf(`invalid Golang environment variable: "%s"`, line)
		}
		array = append(array, []string{gstr.Trim(match[1]), gstr.Trim(match[2])})
	}
	tw := tablewriter.NewWriter(buffer)
	//cmd := exec.Command("cmd")
	//in := bytes.NewBuffer(nil)
	//
	//cmd.Stdin = in
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//
	//go func() {
	//	// start stop restart
	//	in.WriteString("go env\n") //写入你的命令，可以有多行，"\n"表示回车
	//}()
	//if err := cmd.Start(); err != nil {
	//	log.Println(err)
	//}
	//
	//if err := cmd.Wait(); err != nil {
	//	log.Println(err)
	//}
	//var result []string
	//log.Println(string(out.Bytes()))
	//json.Unmarshal(out.Bytes(), &result)
	//log.Println(result)
	return nil
}
