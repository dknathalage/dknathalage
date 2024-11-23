package invoice

type ContactInfo struct {
	ID           int
	Name         string
	BusinessName string
	ABN          string
	Address      string
	Email        string
	Phone        string
}

type PaymentDetails struct {
	ID int
}

type Invoice struct {
	ID          int
	From        ContactInfo
	Participant ContactInfo
	To          ContactInfo
}
