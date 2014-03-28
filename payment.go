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
	Type        string
	Client      *string
	CardType    string `json: "card_type"`
	Code        string
	Account     string
	Country     string
	ExpireMonth int `json: "expire_month"`
	ExpireYear  int `json: "expire_year"`
	Holder      string
	CardHolder  *string `json: "card_holder"`
	Last4       string
	Created     string  `json: "created_at"`
	Updated     string  `json: "updated_at"`
	AppId       *string `json: "app_id"`
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
