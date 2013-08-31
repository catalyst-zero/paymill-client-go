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

func (t *testSuite) TestCreatePayment() {
  token := t.Config.Fixtures.Payment.Token
  c := NewClient(t.Config.Api.Token)

  p := c.CreatePayment(token, nil)

  t.Not(t.Equal(p.Id, ""))
  t.Equal(p.PaymentType(), CreditCard)
}
