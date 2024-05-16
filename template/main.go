package template

func Main(name string) string {
	return `package main

import "` + name + `/boot/serve"

//	@title			Swagger ` + name + ` API
//	@version		1.0
//	@description	` + name + ` project

//	@contact.name	antgo
//	@contact.url	https://github.com/small-ek/antgo
//	@contact.email	56494565@qq.com

// @host		127.0.0.1:9001
// @BasePath	/api
func main() {
	serve.LoadSrv()
}

`
}
