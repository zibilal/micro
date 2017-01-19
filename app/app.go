package app

import (
	"fmt"
	"github.com/maps90/librarian"
	"github.com/mataharimall/micro/api"
)

const BASE_PATH = "$GOPATH/src/github.com/mataharimall/micro"

func Construct() (err error) {
	err = initConfig(BASE_PATH)
	err = InitApp()
	err = initRouter()
	fmt.Print("a")
	return
}

func InitApp() error {
	librarian.Set("loket", func() (interface{}, error) {
		return api.NewLoketApi("loket")
	})
	return nil
}
