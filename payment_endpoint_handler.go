package paymill

import (
  "fmt"
  "net/url"
)

func (c *ApiClient) CreatePayment(token string, client *string) (*Payment, error) {
  values := url.Values{}
  values.Add("token", token)

  if client != nil {
    values.Add("client", *client)
  }

  resp, body := c.doRequest("payments", "POST", nil, values)

  r, err := NewPaymentResponse(resp, body)

  return r.Data, err
}

func (c *ApiClient) PaymentDetails(id string) (*Payment, error) {
  resource := fmt.Sprintf("payments/%s", id)

  resp, body := c.doRequest(resource, "GET", nil, nil)

  r, err := NewPaymentResponse(resp, body)

  return r.Data, err
}

func (c *ApiClient) ListPayments(order string, filter string) (payments []Payment, err error) {
  values := url.Values{}
  if !Empty(order) {
    values.Add("order", order)
  }

  if !Empty(filter) {
    values.Add("filter", filter)
  }

  resource := "payments"

  resp, body := c.doRequest(resource, "GET", values, nil)

  r, err := NewListPaymentResponse(resp, body)

  return r.Data, err
}

func (c *ApiClient) DeletePayment(id string) (ok bool, err error) {
  resource := fmt.Sprintf("payments/%s", id)

  resp, body := c.doRequest(resource, "DELETE", nil, nil)

  _, err = NewDeletePaymentResponse(resp, body)

  return (err == nil), err
}
