package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/maps90/librarian"
	"github.com/mataharimall/micro/api"
	"github.com/mataharimall/micro/helper"
	"net/http"
	"fmt"
)

type EventsList struct {
	Request  interface{}
	Response interface{}
}

func GetEventList(c echo.Context) error {
	loket, ok := librarian.Get("loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")

	}

	loket.GetAuth().Post("/v3/event", "form", "")
	var m map[string]interface{}
	json.Unmarshal([]byte(loket.Body), &m)

	return helper.BuildJSON(c, m)

}

func GetEventSchedule(c echo.Context) error {
	loket, ok := librarian.Get("loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")

	}
	url := fmt.Sprintf(`/v3/shedule/%s`, c.Param("schedule_id"))
	loket.GetAuth().Post(url, "form", "")
	var m map[string]interface{}
	json.Unmarshal([]byte(loket.Body), &m)

	return helper.BuildJSON(c, m)
}
