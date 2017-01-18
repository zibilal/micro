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
	r := &eventsList{}

	if err := c.Bind(r.Request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loket := container.Get("api.loket").(api.Loket)
	loket.GetAuth().Post("/v3/events", "form", "")

	return c.JSON(http.StatusOK, loket.Body)
}
