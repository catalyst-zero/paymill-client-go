package paymill

import (
  "strings"
  "net/http"
)

func Empty(s string) bool {
  return strings.Trim(s, " ") == ""
}

func IsError(r *http.Response) (bool) {
  return r.StatusCode >= 400
}
