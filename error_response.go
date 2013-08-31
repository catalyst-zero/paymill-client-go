package paymill

import (
  _ "fmt"
  "net/http"
  "encoding/json"
  "errors"
)

type ErrorResponse struct {
  Error string
  Exception string
}

func NewErrorResponse(resp *http.Response, body []byte) (error) {
  var r ErrorResponse
  err := json.Unmarshal(body, &r)
  if err != nil {
    panic(err)
  }

  return errors.New(r.Error)
}
