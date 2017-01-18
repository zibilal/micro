package loket

import (
	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
	"github.com/mataharimall/micro-api/commons"
	_ "github.com/mataharimall/micro-api/middlewares"
	"net/http"
	"github.com/mataharimall/micro-api/apps/loket/controller"
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

	e.Post("/loket/invoice/create", controller.CreateInvoice)
	return e
}
