package paymill

import (
  "net/http"
  "encoding/json"
)

type PaymentResponse struct {
  Data *Payment
  Mode string
}

func NewPaymentResponse(resp *http.Response, body []byte) (r *PaymentResponse, e error) {
  err := json.Unmarshal(body, &r)
  if err != nil {
    panic(err)
  }

  if (resp.StatusCode >= 400) {
    e = NewErrorResponse(resp, body)
  }

  return r, e
}
