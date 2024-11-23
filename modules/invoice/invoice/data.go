package invoice

import "time"

type From struct {
	Name    string
	ABN     string
	Address Address
	Email   string
	Phone   string
}

type BillTo struct {
	Name    string
	Company string
	Address Address
	Email   string
	Phone   string
}

type NDISParticipant struct {
	Name    string
	NDISRef string
}

type Address struct {
	Street     string
	PostalCode string
	City       string
	State      string
}

type InvoiceItem struct {
	Item        int
	Description string
	Date        time.Time
	Quantity    int
	UnitPrice   float64
	Total       float64
}

type PaymentMethod struct {
	BankName      string
	BSBNumber     string
	AccountNumber string
}

type Invoice struct {
	From            From
	InvoiceNumber   string
	InvoiceDate     time.Time
	NDISParticipant NDISParticipant
	BillTo          BillTo
	InvoiceSummary  []InvoiceItem
	PaymentMethod   PaymentMethod
}
