package handlers

import (
	"github.com/labstack/echo"
	"github.com/mataharimall/micro/api"
	"github.com/mataharimall/micro/container"
	"github.com/mataharimall/micro/helpers"
	"net/http"
	"encoding/json"
)

type EventsList struct {
	Request interface{}
	Response interface{}
}

func GetEventList(c echo.Context) error {
	loket, ok := container.Get("api.loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")

	}

	loket.GetAuth().Post("/v3/event", "form", "")
	var m map[string]interface{}
	json.Unmarshal([]byte(loket.Body), &m)

	return helpers.BuildJSON(c, m)

}
