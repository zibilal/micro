package app

import (
	"fmt"
	"github.com/labstack/gommon/color"
	config "github.com/spf13/viper"
)

func initConfig(basepath string) (err error) {
	config.SetConfigName("app")
	config.AddConfigPath(basepath)
	if err := config.ReadInConfig(); err != nil {
		return fmt.Errorf("%s: %s", color.Red("ERROR"), color.Yellow("config files not found."))
	}
	return
}
