package helpers

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
	"time"
	"github.com/smartystreets/assertions/should"
	"encoding/json"
)

type TypeStruct struct {
	FirstName string
	LastName string
	BirthDate time.Time
}

type TypeStruct2 struct {
	Category string `json:"cat"`
	ParentId uint `json:"parent_id"`
	Id uint
}

type TypeStruct3 struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	BirthDate time.Time `json:"dob"`
	Category TypeStruct2 `json:"category"`
}

type TypeStruct4 struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	BirthDate string `json:"dob"`
	Category TypeStruct2 `json:"category"`
}

func TestFetchField(t *testing.T) {
	convey.Convey("FetchField test", t, func() {
		cat := TypeStruct2{
			"123", 15, 17,
		}
		val := TypeStruct3{
			"Bilal", "Muhammad", time.Date(1977, time.December, 11, 0, 0, 0, 0, time.UTC),
			cat,
		}

		result, _, err := FetchField("json", "parent_id", val)
		convey.So(err, should.BeNil)
		convey.So(result, should.NotBeNil)
		//convey.Println("Result", result, "Err", err)
	})
}

func TestWalkinMap(t *testing.T) {
	convey.Convey("WalkinMap test", t, func() {
		jsonStr := `{
				    "status": "success",
				    "data": {
					"info_invoice": {
					    "firstname": "Noer",
					    "lastname": "Cholis",
					    "email": "poetnoer@gmail.com",
					    "telephone": null,
					    "identity_id": null,
					    "invoice_date": "02 Nov 2015 15:59",
					    "invoice_expired": "02 Nov 2015 16:14",
					    "invoice_status": "BOOKED",
					    "payment_type": "Cash Deposit"
					},
					"info_summary": {
					    "item-1": {
						"ticket_type": "VIP",
						"ticket_price": "2.000.000",
						"ticket_quantity": "2",
						"ticket_total": "4.000.000"
					    },
					    "item-2": {
						"ticket_type": "GA",
						"ticket_price": "1.000.000",
						"ticket_quantity": "2",
						"ticket_total": "2.000.000"
					    },
					    "ticket_totalx": "6.000.000"
					},
					"invoice_code": "YEN9F5RG"
				    },
				    "message": "Invoice created succesfully"
				}`
		data := map[string]interface{}{}
		err := json.Unmarshal([]byte(jsonStr), &data)

		convey.So(err, should.BeNil)

		convey.Println("Data", data)
		result := TypeStruct4{}
		err = WalkinMap(data, &result)

		convey.So(err, should.BeNil)
		convey.Println("Result", result)
	})
}

func TestMap2Struct(t *testing.T) {
	convey.Convey("Test Map to struct completion", t, func() {
		result1 := &TypeStruct {}
		Map2Struct(map[string]interface{}{
			"FirstName": "Bilal",
			"LastName": "Muhammad",
			"BirthDate": time.Date(2016, time.August, 29, 9, 0, 0, 0, time.UTC),
		}, "app", result1)

		convey.So(result1.FirstName, should.Equal, "Bilal")
		convey.So(result1.LastName, should.Equal, "Muhammad")
		convey.So(result1.BirthDate.String(), should.Equal, time.Date(2016, time.August, 29, 9, 0, 0, 0, time.UTC).String())
	})
}
