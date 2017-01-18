package helpers

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/mataharimall/api-seller/components"
	config "github.com/spf13/viper"
	"net/http"
	"reflect"
	"time"
)

func BuildResponse(c echo.Context, startTime time.Time, rs interface{}, total interface{}, err error) error {
	reqId := RandomString(20)
	rsLen := rs == nil
	val := reflect.ValueOf(rs)
	if val.Kind() == reflect.Ptr && val.IsNil() || rsLen {
		rs = make([]string, 0)
	}
	if err != nil {
		c.Error(err)
		return err
	} else {

		response := map[string]interface{}{
			"code":      http.StatusOK,
			"requestId": reqId,
			"results":   rs,
		}
		if total != nil {
			response["total"] = total
		}
		if config.GetBool("debug") {
			stop := time.Now()
			ss := uint32(stop.Sub(startTime) / time.Millisecond)
			exec_time := fmt.Sprintf("%dms", ss)
			response["execution_time"] = exec_time
		}

		//logger
		component.SaveLog(c, response)

		return c.JSON(http.StatusOK, response)
	}
}
