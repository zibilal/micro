package controller

import (
	"github.com/mataharimall/micro-api/commons"
	"github.com/labstack/echo"
	"github.com/mataharimall/micro-api/services"
	"github.com/mataharimall/micro-api/helpers"
	"time"
)

var CreateInvoice commons.EchoControllerFunc = func(c echo.Context) error {

	startTime := time.Now()

	option := &services.CreateInvoiceRequest{}

	if err := c.Bind(option); err != nil {
		return helpers.BuildResponse(c, startTime, nil, nil, err)
	}

}