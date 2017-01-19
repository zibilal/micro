package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/mataharimall/micro/api"
	"encoding/json"
	"fmt"
	"github.com/mataharimall/micro/helper"
	"github.com/maps90/librarian"
)

type CreateInvoiceRequestResponse struct {
	Request CreateInvoiceRequest
	Response interface{}
}

type CreateInvoiceRequest struct {
	Data Data `json:"data"`
}

type Data struct {
	Tickets        []Ticket `json:"tickets"`
	Attendee       Attendee `json:"attendee"`
	OrderId        string   `json:"order_id"`
	ExpirationType string   `json:"expiration_type"`
	Notes          string   `json:"notes"`
}

type Ticket struct {
	IdTicket string `json:"id_ticket"`
	Quantity string `json:"qty"`
}

type Attendee struct {
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	IdentityId string `json:"identity_id"`
	Dob        string `json:"dob"`
	Gender     string `json:"gender"`
	Email      string `json:"email"`
	Telephone  string `json:"telephone"`
}

func CreateInvoice(c echo.Context) error {
	r := CreateInvoiceRequestResponse{}

	if err := c.Bind(&r.Request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loket, ok := librarian.Get("api.loket").(*api.Loket)

	fmt.Printf("Type %T\n", loket)
	fmt.Printf("Value %v\n", loket)
	if !ok {
		fmt.Println("Internal server error")
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	fmt.Printf("Request %v\n", r.Request)

	jbyte, err := json.Marshal(r.Request)
	str := string(jbyte)
	fmt.Println("String", str)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	loket.GetAuth().Post("/v3/invoice/create", "json", str)
	var m map[string]interface{}
	json.Unmarshal([]byte(loket.Body), &m)

	return helper.BuildJSON(c, m)
}