package services

import (
	"errors"
	"github.com/mataharimall/micro-api/commons"
	"github.com/mataharimall/micro-api/helpers"
	"time"
)

type LoketInvoice struct {
	FirstName      string    `json:"firstname"`
	LastName       string    `json:"lastname"`
	IdentityId     string    `json:"identity_id"`
	Dob            time.Time `json:"dob"`
	Gender         string    `json:"gender"`
	Email          string    `json:"email"`
	Telephone      string    `json:"telephone"`
	OrderId        string    `json:"order_id"`
	ExpirationType string    `json:"expiration_type"`
	Notes          string    `json:"notes"`
	Tickets        []Ticket  `json:"tickets"`
}

type Ticket struct {
	IdTicket    string `json:"id_ticket"`
	Quantity    string `json:"qty"`
	TicketPrice string `json:"ticket_price,omitempty"`
	TicketType  string `json:"ticket_type,omitempty"`
}

func NewLoketInvoice() *LoketInvoice {
	return &LoketInvoice{}
}

func (i *LoketInvoice) PersitenceName() string {
	return "invoice"
}

func (i *LoketInvoice) CreateInvoice(request CreateInvoiceRequest) (interface{}, error) {
	// Copy data for convenience

	i.FirstName = request.Data.Attendee.FirstName
	i.LastName = request.Data.Attendee.LastName
	i.IdentityId = request.Data.Attendee.IdentityId
	i.Dob, _ = time.Parse(commons.YMD_FORMAT, request.Data.Attendee.Dob)
	i.Gender = request.Data.Attendee.Gender
	i.Email = request.Data.Attendee.Email
	i.Telephone = request.Data.Attendee.Telephone
	i.OrderId = request.Data.OrderId
	i.ExpirationType = request.Data.ExpirationType
	i.Notes = request.Data.Notes
	i.Tickets = request.Data.Tickets

	return nil, nil

}

func (i *LoketInvoice) Request(data ...interface{}) (interface{}, error) {
	var returnData bool
	for _, d := range data {
		if d != nil {
			switch v := d.(type) {
			case []Ticket:
				i.Tickets = v
				returnData = true
			case map[string]interface{}:
				helpers.Map2Struct(v, "json", i)
				returnData = true
			default:
				return nil, errors.New("Wrong argument types...")
			}
		}
	}

	if returnData {

		request := struct {
			Data interface{} `json:"data"`
		}{
			Data: struct {
				Tickets        []Ticket    `json:"tickets"`
				Attendee       interface{} `json:"attendee"`
				OrderId        string      `json:"order_id"`
				ExpirationType string      `json:"expiration_type"`
				Notes          string      `json:"notes"`
			}{
				Tickets: i.Tickets,
				Attendee: struct {
					FirstName  string `json:"first_name"`
					LastName   string `json:"last_name"`
					IdentityId string `json:"identity_id"`
					Dob        string `json:"dob"`
					Email      string `json:"email"`
					Telephone  string `json:"telephone"`
				}{
					FirstName:  i.FirstName,
					LastName:   i.LastName,
					IdentityId: i.IdentityId,
					Dob:        i.Dob.Format(commons.YMD_FORMAT),
					Email:      i.Email,
					Telephone:  i.Telephone,
				},
				OrderId:        i.OrderId,
				ExpirationType: i.ExpirationType,
				Notes:          i.Notes,
			},
		}

		return request, nil
	}

	return i, nil
}

func (i *LoketInvoice) Response(data interface{}) (interface{}, error) {
	if data == nil {
		switch  vmap := data.(type) {
		case map[string]interface{}:
			err := helpers.WalkinMap(vmap, i)

			if err != nil {
				return nil, err
			}

		default:
			return nil, errors.New("Unrecognized data type")
		}
	}
	return nil, errors.New("Invalid state")
}
