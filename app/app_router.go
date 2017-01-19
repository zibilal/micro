package app

import (
	"github.com/labstack/echo"
	"github.com/mataharimall/micro/handler"
	"github.com/mataharimall/micro/middleware"
	"github.com/labstack/echo/engine/standard"
	c "github.com/spf13/viper"
	"github.com/facebookgo/grace/gracehttp"
)

func initRouter() error {
	e := echo.New()
	e.SetDebug(c.GetBool("app.debug"))
	e.Use(middleware.Logger())

	e.SetBinder(middleware.AppBinder{})
	e.SetHTTPErrorHandler(middleware.AppHttpErrorHandler)

	e.Get("/loket/event", handler.GetEventList)
	e.Get("/loket/schedule/:schedule_id", handler.GetEventSchedule)
	e.Post("/loket/invoice/paid/:code", handler.PostInvoiceStatus)
	e.Post("/loket/invoice/create", handler.CreateInvoice)
	e.Post("/loket/invoice/list/:invoice_code/attendee", handler.FetchInvoiceListAttendee)

	std := standard.New(":" + c.GetString("app.port"))
	std.SetHandler(e)

	err := gracehttp.Serve(std.Server)
	return err
}
