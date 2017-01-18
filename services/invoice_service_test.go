package services

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
	"github.com/smartystreets/assertions/should"
	"encoding/json"
	"github.com/mataharimall/micro-api/commons"
	"time"
)

func TestLoketInvoice_Request(t *testing.T) {
	convey.Convey("Invoice Request", t, func() {

		data := map[string]interface{}{
			"firstname": "Bilal",
			"lastname": "Muhammad",
			"identity_id": "ktp43344344789",
			"dop": "1945-08-17",
			"gender": "1",
			"email": "bmuhamm@gmail.com",
			"telephone": "021-76911788",
			"order_id": "INV123Until255Char",
			"expiration_type": "2",
		}

		tickets := []Ticket{
			{
				"12341234",
				"12",
			},
			{
				"321324321",
				"24",
			},
		}

		loketi := &LoketInvoice{}
		result, err := loketi.Request(data, tickets)

		convey.So(len(loketi.Tickets), should.Equal, 2)
		convey.So(loketi.FirstName, should.Equal, "Bilal")
		convey.So(loketi.OrderId, should.Equal, "INV123Until255Char")

		convey.So(err, should.BeNil)
		convey.So(result, should.NotBeNil)

		jbyte, err := json.MarshalIndent(result, "", "\t")

		convey.So(err, should.BeNil)
		convey.So(len(jbyte), should.NotBeEmpty)

		//convey.Println(string(jbyte))
	})

	convey.Convey("Invoice request if no input", t, func() {
		timeDob, _ := time.Parse(commons.YMD_FORMAT, "1945-08-17")
		loketi := &LoketInvoice{
			FirstName: "Babang",
			LastName: "Jethro",
			IdentityId: "ktp43344344789",
			Dob: timeDob,
			Gender: "1",
			Email: "bmuhamm@gmail.com",
			Telephone: "021-788978",
			OrderId: "INV123Until255Char",
			ExpirationType: "2",
			Tickets: [] Ticket{
				{
					"12341234",
					"12",
				},
				{
					"321324321",
					"24",
				},
			},
		}

		var iTest commons.Service = loketi
		result, err := iTest.Request(nil)
		convey.So(err, should.BeNil)

		jbyte, err := json.MarshalIndent(result, "", "\t")
		convey.So(len(jbyte), should.NotBeEmpty)

		//convey.Println(string(jbyte))
	})
}
