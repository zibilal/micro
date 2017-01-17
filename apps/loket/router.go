package loket

import (
	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
	"github.com/mataharimall/micro-api/commons"
	_ "github.com/mataharimall/micro-api/middlewares"
	"net/http"
)

type LoketRoute struct{

}

func init() {
	commons.RouterManager.Register("route.loket", &LoketRoute{})
}

func (l *LoketRoute) SetRoute(e *echo.Echo) *echo.Echo {
	e.Use(em.Logger())
	e.Post("/loket/event", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Post("/loket/invoice/create", func(c echo.Context) error {
		// 1. create invoice
		// 2. save to mongo
		return c.String(http.StatusOK, "Hello, World!")
	})
	return e
}
