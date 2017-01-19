package handlers

import (
	"github.com/labstack/echo"
	"github.com/mataharimall/micro/container"
	"github.com/mataharimall/micro/api"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/mataharimall/micro/helpers"
)

var logger = logrus.New()

type InvoiceListAttendee struct {
	Request ListResponse
}

type ListResponse struct {
	InvoiceCode string `json:"invoice_code"`
}

func FetchInvoiceListAttendee(c echo.Context) error {

	r := InvoiceListAttendee{}

	if err := c.Bind(&r.Request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loket, ok := container.Get("api.loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	if r.Request.InvoiceCode == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Invoice code is required")
	}

	url := fmt.Sprintf("/v1/invoice/%s/attendee", r.Request.InvoiceCode)

	loket.GetAuth().Post(url, "form", "")
	var m map[string]interface{}
	err := json.Unmarshal([]byte(loket.Body), &m)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"Endpoint": "FetchInvoiceListAttendee",
		}).Debug("Error:", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to decode server response")
	}

	return helpers.BuildJSON(c, m)
}