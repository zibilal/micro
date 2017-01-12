package loket

import (
	app "github.com/mataharimall/micro-api"
	. "github.com/mataharimall/micro-api/commons/idata/assertion"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func init() {
	app.InitConfig()
}

func TestGetAuth(t *testing.T) {
	Convey("Testing Loket API", t, func() {
		Convey("should return token", func() {
			l := New().GetAuth()
			byt := []byte(l.Body)
			So(byt, ShouldBeJSONAndHave, "status", "success")
			So(byt, ShouldBeJSONAndHave, "code", "200")
		})
	})
}
