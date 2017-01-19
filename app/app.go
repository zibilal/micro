package app

import (
	"github.com/maps90/librarian"
	d "github.com/maps90/librarian/datasource"
	"github.com/maps90/librarian/cache"
	"github.com/mataharimall/micro/api"
	c "github.com/spf13/viper"
)

const BASE_PATH = "$GOPATH/src/github.com/mataharimall/micro"

func Construct() (err error) {
	err = initConfig(BASE_PATH)
	err = InitApp()
	err = initRouter()
	return
}

func InitApp() error {
	librarian.Set("loket", func() (interface{}, error) {
		return api.NewLoketApi("loket")
	})
	librarian.Set("mysql.master", func() (interface{}, error) {
		return d.NewDatasourceFactory("mysqlaccess", "", c.GetString("mysql.master"))
	})
	librarian.Set("mysql.slave", func() (interface{}, error) {
		return d.NewDatasourceFactory("mysqlaccess", "", c.GetString("mysql.slave"))
	})
	librarian.Set("redis.master", func() (interface{}, error) {
		return cache.NewCacheStorageFactory("redis", c.GetString("redis.master"), "0")
	})
	librarian.Set("redis.slave", func() (interface{}, error) {
		return cache.NewCacheStorageFactory("redis", c.GetString("redis.slave"), "0")
	})

	return nil
}
