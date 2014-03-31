package paymill

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Empty(s string) bool {
	return strings.Trim(s, " ") == ""
}

func IsError(r *http.Response) bool {
	return r.StatusCode >= 400
}

func intToTime(i int64) time.Time {
	return time.Unix(i, 0)
}

func toTime(s string) time.Time {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return intToTime(i)
}
