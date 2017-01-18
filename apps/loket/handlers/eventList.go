package handlers

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/mataharimall/micro/container"
	"github.com/mataharimall/micro/api"
)

type eventsList struct {
	Request  CommonRequest
	Response struct{}
}

type CommonRequest struct {
	Page    string
	Limit   string
	SortBy  string
	OrderBy string
}

func GetEventList(c echo.Context) error {
	r := new(eventsList)

	if err := c.Bind(&r.Request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loket, ok := container.Get("api.loket").(*api.Loket)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")

	}

	loket.GetAuth().Post("/v3/event", "form", "")
	return c.JSON(http.StatusOK, loket.Body)

}
