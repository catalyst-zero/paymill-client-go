package paymill

import (
	"encoding/json"
	"net/http"
)

type ListPaymentResponse struct {
	Data      []Payment
	DataCount int
	Mode      string
}

func NewListPaymentResponse(resp *http.Response, body []byte) (r *ListPaymentResponse, e error) {
	err := json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}

	if IsError(resp) {
		e = NewErrorResponse(resp, body)
	}

	return r, e
}
