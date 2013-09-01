package paymill

import (
  "strings"
  "net/http"
  "time"
  "strconv"
)

func Empty(s string) bool {
  return strings.Trim(s, " ") == ""
}

func IsError(r *http.Response) (bool) {
  return r.StatusCode >= 400
}

func toTime(s string) time.Time {
  i, err := strconv.ParseInt(s, 10, 64)
  if err != nil {
    panic(err)
  }

  return time.Unix(i, 0)
}
