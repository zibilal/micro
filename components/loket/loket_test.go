package loket

import (
	app "github.com/mataharimall/micro-api"
	. "github.com/mataharimall/micro-api/commons/idata/assertion"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)

func init() {
	app.InitConfig()
}

type Events struct {
	Status string `json:"status"`
	Data []*struct {
		IdEvent string `json:"id_event"`
	} `json:"data"`
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

func TestGetEvents(t *testing.T) {
	Convey("should retun event list",t, func() {
		l := New().GetAuth()
		e := new(Events)
		evt := l.Post("v3", "event",  fmt.Sprintf(`{"token": "%s"}`, l.Token))
		evt.SetStruct(e)
		So(e.Status, ShouldEqual, "success")
	})
}