package main

import (
	"fmt"
	"github.com/labstack/gommon/color"
	config "github.com/spf13/viper"
	"log"
	"os"
	"runtime"
)

const (
	BASE_PATH = "$GOPATH/src/github.com/mataharimall/micro-api"
	APP_VER   = "v1"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := InitConfig(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func InitConfig() (err error) {
	config.SetConfigName("app")
	config.AddConfigPath(BASE_PATH + "/config")
	if err := config.ReadInConfig(); err != nil {
		log.Println(err)
		return fmt.Errorf("%s: %s", color.Red("ERROR"), color.Yellow("config files not found."))
	}
	return
}

func main() {
}
