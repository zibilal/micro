package middleware

import (
	"github.com/labstack/echo"
	"log"
	"net/http"

	cm "github.com/mataharimall/micro-api/commons"
)

type AppError struct {
	Code         int    `json:"code"`
	requestId    string `json:"requestId"`
	errorMessage string `json:"errorMessage"`
}

func (e *AppError) Error() string {
	return e.errorMessage
}

var AppHttpErrorHandler = func(err error, c echo.Context) {
	if he, ok := err.(*AppError); ok {
		if !c.Response().Committed() {
			c.JSON(he.Code, he)
		}
		log.Println(he)
		return
	}

	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Error()
	}

	reqId := cm.RandomString(20)

	if !c.Response().Committed() {
		// replace all error into json
		response := map[string]interface{}{
			"code":         code,
			"requestId":    reqId,
			"errorMessage": msg,
		}

		c.JSON(code, response)
	}
	log.Println(err)
}
