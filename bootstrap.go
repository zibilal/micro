package micro

import (
	"fmt"
	"os"

	"github.com/labstack/gommon/color"
	"github.com/mataharimall/micro/apps/loket"
	"github.com/mataharimall/micro/service"
	config "github.com/spf13/viper"
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
}

func initService() {
	service.ServiceManager.Register("route.loket", &loket.LoketRoute{})
	service.ServiceManager.Init()
}

func Init() error {
	if err := initConfig(); err != nil {
		return err
		os.Exit(1)
	}
	initContainer()
	initService()
	return nil
}
