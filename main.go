package main

import (
	"flag"
	"fmt"
	"strings"
)

var name = flag.Int("name", 1, "学生的姓名")

var major string

func init() {
	const (
		defaultMajor = "计算机"
		usage        = "学生的专业"
	)
	flag.StringVar(&major, "major", "1212", usage)
	flag.StringVar(&major, "m", "122112", usage+" (简写)")
}

type Classmates []string

func (i *Classmates) String() string {
	return fmt.Sprint(*i)
}

func (i *Classmates) Set(value string) error {
	for _, dt := range strings.Split(value, ",") {
		*i = append(*i, dt)
	}
	return nil
}

var mates Classmates

func init() {
	flag.Var(&mates, "class", "逗号分隔的同学列表")
}

func main() {
	flag.Parse()
	fmt.Println("name ", *name)
	fmt.Println("major ", major)
	fmt.Println("classmates ", mates)
}