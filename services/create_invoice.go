package services

type CreateInvoiceRequest struct {
	Data DataRequest `json:"data"`
}

type DataRequest struct {
	Attendee       InvoiceAttendee `json:"attendee"`
	Tickets        []Ticket        `json:"tickets"`
	OrderId        string          `json:"order_id"`
	ExpirationType string          `json:"expiration_type"`
	Notes          string          `json:"notes"`
}

type InvoiceAttendee struct {
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	IdentityId string `json:"identity_id"`
	Dob        string `json:"dob"`
	Gender     string `json:"gender"`
	Email      string `json:"email"`
	Telephone  string `json:"telephone"`
}
