package main

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/gommon/color"
	"github.com/mataharimall/micro-api/routes"
	config "github.com/spf13/viper"
	"log"
	"os"
	"runtime"
)

const BASE_PATH = "$GOPATH/src/github.com/mataharimall/micro-api"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := InitConfig(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func InitConfig() (err error) {
	config.SetConfigName("config")
	config.AddConfigPath(BASE_PATH)
	if err := config.ReadInConfig(); err != nil {
		log.Println(err)
		return fmt.Errorf("%s: %s", color.Red("ERROR"), color.Yellow("config files not found."))
	}
	return
}

func main() {
	r := routes.SetRoute()
	std := standard.New(":" + config.GetString("port"))
	std.SetHandler(r)
	gracehttp.Serve(std.Server)
}
