package paymill

import (
	_ "fmt"
	"time"
)

type PaymentType int

const (
	CreditCard PaymentType = iota
	DebitCard
)

// Object for direct debit payments
type Payment struct {
	Id          string
	Type        string `json:"payment_type"`
	Client      *string
	CardType    string
	Code        string
	Account     string
	Country     string
	ExpireMonth int
	ExpireYear  int
	Holder      string
	CardHolder  *string
	Last4       string
	Created     string `json: "created_at"`
	Updated     string `json: "updated_at"`
	AppId       *string
}

func (p *Payment) PaymentType() PaymentType {
	return CreditCard
}

func (p *Payment) CreatedAt() time.Time {
	return toTime(p.Created)
}

func (p *Payment) UpdatedAt() time.Time {
	return toTime(p.Updated)
}
