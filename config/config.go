package config

import (
	"fmt"
	"github.com/labstack/gommon/color"
	config "github.com/spf13/viper"
	"log"
)

const BASE_PATH = "$GOPATH/src/github.com/mataharimall/micro-api/config"

func Init() (err error) {
	config.SetConfigName("app")
	config.AddConfigPath(BASE_PATH)
	if err := config.ReadInConfig(); err != nil {
		log.Println(err)
		return fmt.Errorf("%s: %s", color.Red("ERROR"), color.Yellow("config files not found."))
	}
	return
}
