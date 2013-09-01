package paymill

import (
  _ "fmt"
  "github.com/remogatto/prettytest"
  "testing"
  "io/ioutil"
  "launchpad.net/goyaml"
)

type TestApi struct {
  Token string `yaml: "token"`
}

type TestPayment struct {
  Token string
  Id string
}

type TestFixtures struct {
  Payment TestPayment
}

type TestConfig struct {
  Api TestApi
  Fixtures TestFixtures
}

// Start of setup
type testSuite struct {
  prettytest.Suite
  Config TestConfig
}

func TestRunner(t *testing.T) {
  b, err := ioutil.ReadFile("support/test.yml")
  if err != nil {
    panic(err)
  }

  var c TestConfig

  goyaml.Unmarshal(b, &c)

  ts := new(testSuite)
  ts.Config = c

  prettytest.RunWithFormatter(
      t,
      new(prettytest.TDDFormatter),
      ts,
  )
}
// End of setup

func (t *testSuite) TestCreateCreditCardPayment() {
  token := t.Config.Fixtures.Payment.Token
  c := NewClient(t.Config.Api.Token)

  p, err := c.CreatePayment(token, nil)

  t.Equal(err, nil)

  t.Not(t.Equal(p.Id, ""))
  t.Equal(p.PaymentType(), CreditCard)
}

func (t *testSuite) TestPaymentDetailsForNonExistingPayment() {
  id := "not_found"
  c := NewClient(t.Config.Api.Token)

  _, err := c.PaymentDetails(id)

  t.Equal("Payment not Found", err.Error())
}

func (t *testSuite) TestListPayments() {
  c := NewClient(t.Config.Api.Token)

  p, err := c.ListPayments("count", "card_type=visa")

  t.Equal(err, nil)

  t.Equal(len(p), 20)
}
