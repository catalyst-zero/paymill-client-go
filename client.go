package paymill

import (
  "fmt"
  "strings"
  "net/http"
  "net/url"
  "io/ioutil"
)

const APIBase string = "https://api.paymill.com"
const APIVersion string = "v2"

type Client struct {
  Token string
}

func NewClient(token string) (c *Client) {
  if strings.Trim(token, " ") == "" {
    return nil
  }

  c = &Client{
    Token: token,
  }
  return
}

func BaseUrl() (string) {
  return fmt.Sprintf("%s/%s", APIBase, APIVersion)
}

func UrlFor(entity string) (string) {
  return fmt.Sprintf("%s/%s", BaseUrl(), entity)
}

func (c *Client) CreatePayment(token string, client *string) (*Payment) {
  values := url.Values{}
  values.Add("token", token)

  if client != nil {
    values.Add("client", *client)
  }

  http_client := &http.Client{}

  // This can be wrapped in a method
  var req *http.Request
  req, err := http.NewRequest("POST", UrlFor("payments"), strings.NewReader(values.Encode()))
  if err != nil {
    panic(err)
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.SetBasicAuth(c.Token, "")

  var resp *http.Response
  resp, err = http_client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  return NewPaymentResponse(body).Data
}
