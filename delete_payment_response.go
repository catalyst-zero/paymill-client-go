package paymill

import (
	"encoding/json"
	"net/http"
)

type DeletePaymentResponse struct {
	Data []interface{}
}

func NewDeletePaymentResponse(resp *http.Response, body []byte) (r *DeletePaymentResponse, e error) {
	err := json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}

	if IsError(resp) {
		e = NewErrorResponse(resp, body)
	}

	return r, e
}
