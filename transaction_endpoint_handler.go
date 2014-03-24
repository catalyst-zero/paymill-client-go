package paymill

import (
	"net/url"
)

type TransactionPayload struct {
	PaymentId   string
	Amount      string
	Currency    string
	Description string
}

func (c *ApiClient) CreateTransaction(transactionPayload TransactionPayload) (*Transaction, error) {
	values := url.Values{}
	values.Add("payment", transactionPayload.PaymentId)
	values.Add("amount", transactionPayload.Amount)
	values.Add("currency", transactionPayload.Currency)
	values.Add("description", transactionPayload.Description)

	resp, body := c.doRequest("transactions", "POST", nil, values)

	r, err := newTransactionResponse(resp, body)

	return r.Data, err
}
