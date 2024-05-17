package serve

import (
	"embed"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/ant-cli/boot/router"
	"github.com/small-ek/antgo/frame/ant"
	_ "github.com/small-ek/antgo/frame/serve/gin"
)

// LoadSrv Load Api service<加载API服务>
func LoadSrv(f embed.FS) {
	gin.ForceConsoleColor()

	configPath := flag.String("config", "./config/config.toml", "Configuration file path")

	flag.Parse()

	eng := ant.New(*configPath).SetPort("49000").Serve(router.Load(f))
	defer eng.Close()
}
