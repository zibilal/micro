package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	appMiddleware "github.com/mataharimall/micro-api/routes/middleware"
	config "github.com/spf13/viper"
)

func SetRoute() *echo.Echo {
	/*start echo*/
	e := echo.New()
	// Set MiddleWare
	if config.GetBool("debug") {
		e.Use(appMiddleware.Logger())
	}
	e.Use(middleware.Recover(), middleware.Gzip())
	e.SetBinder(appMiddleware.AppBinder{})
	e.SetHTTPErrorHandler(appMiddleware.AppHttpErrorHandler)

	//routing goes here

	return e
}
