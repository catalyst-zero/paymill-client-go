package paymill

import (
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
  ApiClient *ApiClient
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
  ts.ApiClient = NewApiClient(c.Api.Token)

  prettytest.RunWithFormatter(
      t,
      new(prettytest.TDDFormatter),
      ts,
  )
}
// End of setup
