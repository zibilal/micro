package micro

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	c "github.com/spf13/viper"
	"sync"
)

var RouterManager = &router{}

type router struct {
	mux        sync.RWMutex
	collection map[string]interface{}
}

type collection interface {
	SetRoute(*echo.Echo) *echo.Echo
}

func (r *router) Register(name string, col collection) {
	if r.collection == nil {
		r.collection = make(map[string]interface{})
	}
	r.mux.Lock()
	r.collection[name] = col
	r.mux.Unlock()
}

func (r *router) Init() {
	e := echo.New()
	e.SetDebug(c.GetBool("app.debug"))
	fmt.Println("%#v", r)
	for _, v := range r.collection {
		if d, ok := v.(collection); ok {
			e = d.SetRoute(e)
		}
	}
	std := standard.New(":" + c.GetString("app.port"))
	std.SetHandler(e)
	gracehttp.Serve(std.Server)
}
