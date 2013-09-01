package paymill

import (
  "net/url"
)

func (c *ApiClient) CreateClient(email string, description string) (*Client, error) {
  values := url.Values{}
  if !Empty(email) {
    values.Add("email", email)
  }

  if !Empty(description) {
    values.Add("description", description)
  }

  resp, body := c.doRequest("clients", "POST", values)

  r, err := NewClientResponse(resp, body)

  return r.Data, err
}
