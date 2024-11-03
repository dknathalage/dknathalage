package main

import (
	"time"
)

type Invoice struct {
	From            ContactInfo     `json:"from"`
	InvoiceDetails  InvoiceDetails  `json:"invoice_details"`
	NDISParticipant NDISParticipant `json:"ndis_participant"`
	BillTo          ContactInfo     `json:"bill_to"`
	InvoiceSummary  []InvoiceItem   `json:"invoice_summary"`
	PaymentMethod   PaymentMethod   `json:"payment_method"`
	Notes           InvoiceNotes    `json:"notes"`
}

type ContactInfo struct {
	Name    string `json:"name"`
	Company string `json:"company,omitempty"`
	ABN     string `json:"abn,omitempty"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type InvoiceDetails struct {
	InvoiceNumber string    `json:"invoice_number"`
	InvoiceDate   time.Time `json:"invoice_date"`
}

type NDISParticipant struct {
	ParticipantName string `json:"participant_name"`
	NDISRef         string `json:"ndis_ref"`
}

type InvoiceItem struct {
	ItemNumber   int       `json:"item_number"`
	Code         string    `json:"code"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
	Quantity     int       `json:"quantity"`
	UnitPriceAUD float64   `json:"unit_price_aud"`
	TotalAUD     float64   `json:"total_aud"`
	Details      string    `json:"details"`
}

type PaymentMethod struct {
	BankName      string `json:"bank_name"`
	BSBNumber     string `json:"bsb_number"`
	AccountNumber string `json:"account_number"`
}

type InvoiceNotes string
