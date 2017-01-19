package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/maps90/librarian"
	"github.com/mataharimall/micro/api"
	"github.com/mataharimall/micro/helper"
	"net/http"
)

func PostInvoiceStatus(c echo.Context) error {
	loket, ok := librarian.Get("api.loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	url := fmt.Sprintf(`/v3/invoice/%s/paid`, c.Param("code"))
	loket.GetAuth().Post(url, "form", "")
	var m map[string]interface{}
	json.Unmarshal([]byte(loket.Body), &m)

	return helper.BuildJSON(c, m)
}
