package paymill

import (
  "fmt"
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

  resp, body := c.doRequest("clients", "POST", nil, values)

  r, err := NewClientResponse(resp, body)

  return r.Data, err
}

func (c *ApiClient) ClientDetails(id string) (*Client, error) {
  resource := fmt.Sprintf("clients/%s", id)

  resp, body := c.doRequest(resource, "GET", nil, nil)

  r, err := NewClientResponse(resp, body)

  return r.Data, err
}

func (c *ApiClient) ClientUpdate(id string, email string, description string) (*Client, error) {
  values := url.Values{}
  if !Empty(email) {
    values.Add("email", email)
  }

  if !Empty(description) {
    values.Add("description", description)
  }

  resource := fmt.Sprintf("clients/%s", id)

  resp, body := c.doRequest(resource, "PUT", nil, values)

  r, err := NewClientResponse(resp, body)

  return r.Data, err
}

func (c *ApiClient) RemoveClient(id string) (ok bool, err error) {
  resource := fmt.Sprintf("clients/%s", id)

  resp, body := c.doRequest(resource, "DELETE", nil, nil)

  _, err = NewClientResponse(resp, body)

  return (err == nil), err
}

func (c *ApiClient) ListClients(order string, filter map[string]string) (payments []Client, err error) {
  values := url.Values{}
  if !Empty(order) {
    values.Add("order", order)
  }

  for k, v := range filter {
    values.Add(k, v)
  }

  resource := "clients"

  resp, body := c.doRequest(resource, "GET", values, nil)

  r, err := NewListClientsResponse(resp, body)

  return r.Data, err
}
