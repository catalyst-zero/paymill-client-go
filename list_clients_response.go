package paymill

import (
  "net/http"
  "encoding/json"
)

type ListClientsResponse struct {
  Data []Client
  DataCount int
  Mode string
}

func NewListClientsResponse(resp *http.Response, body []byte) (r *ListClientsResponse, e error) {
  err := json.Unmarshal(body, &r)
  if err != nil {
    panic(err)
  }

  if IsError(resp) {
    e = NewErrorResponse(resp, body)
  }

  return r, e
}
