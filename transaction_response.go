package paymill

import (
	"encoding/json"
	"net/http"
)

type TransactionResponse struct {
	Data *Transaction
	Mode string
}

func NewTransactionResponse(resp *http.Response, body []byte) (r *TransactionResponse, e error) {
	err := json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}

	if IsError(resp) {
		e = NewErrorResponse(resp, body)
	}

	return r, e
}
