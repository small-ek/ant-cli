package template

func Main(name string) string {
	return `package main

import "` + name + `/boot/serve"

func main() {
	serve.LoadSrv()
}

`
}
