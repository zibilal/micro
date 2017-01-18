package loket

import (
	"fmt"
	. "github.com/mataharimall/micro-api/helpers/idata/assertion"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/mataharimall/micro-api/config"
	"testing"
)

func init() {
	config.Init()
}

type Events struct {
	Status string `json:"status"`
	Data   []*struct {
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
	Convey("should retun event list", t, func() {
		l := New().GetAuth()
		e := new(Events)
		evt := l.Post("/v3/event", "form", fmt.Sprintf(`{"token": "%s"}`, l.Token))
		evt.SetStruct(e)
		So(e.Status, ShouldEqual, "success")
	})
}

func TestLoket_Configure(t *testing.T) {
	Convey("Configure test", t, func() {
		l := New().Configure(map[string]interface{}{
			"UserName": "testinguser",
			"Password": "testingpassword",
			"ApiKey": "12341234124",
		})
		So("testinguser", ShouldEqual, l.UserName)
		So("testingpassword", ShouldEqual, l.Password)
		So("12341234124", ShouldEqual, l.ApiKey)
	})
}
