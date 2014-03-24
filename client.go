package paymill

import (
	"time"
)

type Client struct {
	Id          string
	Email       string
	Description string
	Created     string `json: "created_at"`
	Updated     string `json: "updated_at"`
	Payment     []string
	/* Subscription []Subscription */
	AppId string
}

func (c *Client) CreatedAt() time.Time {
	return toTime(c.Created)
}

func (c *Client) UpdatedAt() time.Time {
	return toTime(c.Updated)
}
