package paymill

import (
	_ "fmt"
	"time"
)

// Object for direct debit payments
type Payment struct {
	Id          string  `json:"id"`
	Type        string  `json:"type"`
	Client      *string `json:"client"`
	CardType    string  `json:"card_type"`
	Code        string  `json:"code"`
	Account     string  `json:"account"`
	Country     string  `json:"country"`
	ExpireMonth string  `json:"expire_month"`
	ExpireYear  string  `json:"expire_year"`
	Holder      string  `json:"holder"`
	CardHolder  *string `json:"card_holder"`
	Last4       string  `json:"last4"`
	Created     int64   `json:"created_at"`
	Updated     int64   `json:"updated_at"`
	AppId       *string `json:"app_id"`
}

func (p *Payment) CreatedAt() time.Time {
	return intToTime(p.Created)
}

func (p *Payment) UpdatedAt() time.Time {
	return intToTime(p.Updated)
}
