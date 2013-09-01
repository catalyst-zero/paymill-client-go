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

type ApiClient struct {
  Token string
}

func NewApiClient(token string) (c *ApiClient) {
  if strings.Trim(token, " ") == "" {
    return nil
  }

  c = &ApiClient{
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

func (c *ApiClient) doRequest(resource string, method string, data url.Values) (resp *http.Response, body []byte) {
  http_client := &http.Client{}

  // This can be wrapped in a method
  var req *http.Request
  req, err := http.NewRequest(method, UrlFor(resource), strings.NewReader(data.Encode()))
  if err != nil {
    panic(err)
  }

  if method == "POST" {
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  }
  req.SetBasicAuth(c.Token, "")

  resp, err = http_client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()

  body, err = ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err)
  }

  return
}

func (c *ApiClient) CreatePayment(token string, client *string) (*Payment, error) {
  values := url.Values{}
  values.Add("token", token)

  if client != nil {
    values.Add("client", *client)
  }

  resp, body := c.doRequest("payments", "POST", values)

  r, err := NewPaymentResponse(resp, body)

  return r.Data, err
}

func (c *ApiClient) PaymentDetails(id string) (*Payment, error) {
  values := url.Values{}
  resource := fmt.Sprintf("payments/%s", id)

  resp, body := c.doRequest(resource, "GET", values)

  r, err := NewPaymentResponse(resp, body)

  return r.Data, err
}

func (c *ApiClient) ListPayments(order string, filter string) (payments []Payment, err error) {
  values := url.Values{}
  if !Empty(order) {
    values.Add("order", order)
  }

  if !Empty(filter) {
    values.Add("filter", filter)
  }

  resource := "payments"

  resp, body := c.doRequest(resource, "GET", values)

  r, err := NewListPaymentResponse(resp, body)

  return r.Data, err
}

func (c *ApiClient) DeletePayment(id string) (ok bool, err error) {
  values := url.Values{}

  resource := fmt.Sprintf("payments/%s", id)

  resp, body := c.doRequest(resource, "DELETE", values)

  _, err = NewDeletePaymentResponse(resp, body)

  return (err == nil), err
}
