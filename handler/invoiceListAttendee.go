package handler

import (
	"github.com/labstack/echo"
	"github.com/mataharimall/micro/api"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/maps90/librarian"
	"github.com/mataharimall/micro/helper"
)

var logger = logrus.New()

type InvoiceListAttendee struct {
	Request ListResponse
}

type ListResponse struct {
	InvoiceCode string `json:"invoice_code"`
}

func FetchInvoiceListAttendee(c echo.Context) error {
	loket, ok := librarian.Get("loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	url := fmt.Sprintf("/v1/invoice/%s/attendee", c.Param("invoice_code"))

	loket.GetAuth().Post(url, "form", "")
	var m map[string]interface{}
	err := json.Unmarshal([]byte(loket.Body), &m)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"Endpoint": "FetchInvoiceListAttendee",
		}).Debug("Error:", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to decode server response")
	}

	return helper.BuildJSON(c, m)
}