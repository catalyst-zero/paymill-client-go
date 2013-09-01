package paymill

var clientId string

func (t *testSuite) TestCreateClient() {
  c := NewApiClient(t.Config.Api.Token)

  client, err := c.CreateClient("me@example.com", "Foo")

  t.Equal(err, nil)

  t.Not(t.Equal(client.Id, ""))
  t.Equal(client.Email, "me@example.com")
  t.Equal(client.Description, "Foo")
}
