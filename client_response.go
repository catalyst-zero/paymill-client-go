package paymill

import (
  "net/http"
  "encoding/json"
)

type ClientResponse struct {
  Data *Client
}

func NewClientResponse(resp *http.Response, body []byte) (r *ClientResponse, e error) {
  err := json.Unmarshal(body, &r)
  if err != nil {
    panic(err)
  }

  if IsError(resp) {
    e = NewErrorResponse(resp, body)
  }

  return r, e
}
