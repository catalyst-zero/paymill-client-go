package paymill

import (
  "fmt"
  "encoding/json"
)

type PaymentResponse struct {
  Data *Payment
  Mode string
}

func NewPaymentResponse(b []byte) (r *PaymentResponse) {
  err := json.Unmarshal(b, &r)
  if err != nil {
    fmt.Printf("%s", string(b))
    panic(err)
  }

  return r
}
