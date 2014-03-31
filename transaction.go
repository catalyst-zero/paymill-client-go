package paymill

import (
	_ "fmt"
	"time"
)

type Transaction struct {
	Id           string
	Amount       string
	OriginAmount int `json:"origin_amount"`
	Status       string
	Description  *string
	LiveMode     bool `json:"livemode"`
	IsFraud      bool `json:"is_fraud"`
	//Refunds []Refund
	Currency     string
	Created      int64  `json:"created_at"`
	Updated      int64  `json:"updated_at"`
	ResponseCode int    `json:"response_code"`
	ShortId      string `json:"short_id"`
	//Invoices []string
	Payment Payment
	Client  Client
	//Preauthorization Preauthorization
	//Fees []string
	AppId *string `json: "app_id"`
}

func (t *Transaction) CreatedAt() time.Time {
	return intToTime(t.Created)
}

func (t *Transaction) UpdatedAt() time.Time {
	return intToTime(t.Updated)
}
