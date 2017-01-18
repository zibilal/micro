package micro

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/color"
	"github.com/mataharimall/micro/api"
	"github.com/mataharimall/micro/container"
	"github.com/mataharimall/micro/service"
	config "github.com/spf13/viper"

	_ "github.com/mataharimall/micro/apps/loket"
)

const BASE_PATH = "$GOPATH/src/github.com/mataharimall/micro/config"

func initConfig() (err error) {
	config.SetConfigName("app")
	config.AddConfigPath(BASE_PATH)
	if err := config.ReadInConfig(); err != nil {
		return fmt.Errorf("%s: %s", color.Red("ERROR"), color.Yellow("config files not found."))
	}
	return
}

func initContainer() {
	container.Set("api.loket", func() (interface{}, error) {
		return api.NewLoketApi("loket")
	})
}

func Construct() error {
	if err := initConfig(); err != nil {
		return err
		os.Exit(1)
	}
	initContainer()
	service.ServiceManager.Init()
	return nil
}
