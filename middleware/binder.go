package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

const (
	AppJsonHeader = "application/vnd.api+json"
)

type AppBinder struct {
	echo.Binder
}

func (AppBinder) Bind(i interface{}, c echo.Context) (err error) {
	req := c.Request()
	ctype := req.Header().Get(echo.HeaderContentType)
	if req.Body() == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request body can't be empty")
	}
	switch {
	case strings.Contains(ctype, AppJsonHeader):
		ct := strings.Split(ctype, ";")
		if len(ct) > 0 {
			ctype = ct[0]
		}
		if strings.EqualFold(ctype, AppJsonHeader) {
			if err = json.NewDecoder(req.Body()).Decode(i); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		} else {
			c.Error(echo.ErrUnsupportedMediaType)
			return echo.ErrUnsupportedMediaType
		}
	default:
		c.Error(echo.ErrUnsupportedMediaType)
		return echo.ErrUnsupportedMediaType
	}
	return
}
