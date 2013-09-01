package paymill

var clientId string

func (t *testSuite) TestClient() {
  client, err := t.ApiClient.CreateClient("me@example.com", "Foo")

  t.Equal(err, nil)

  clientId = client.Id

  t.Not(t.Equal(client.Id, ""))
  t.Equal(client.Email, "me@example.com")
  t.Equal(client.Description, "Foo")
}

func (t *testSuite) TestClientDetails() {
  client, err := t.ApiClient.ClientDetails(clientId)

  t.Equal(err, nil)

  t.Equal(client.Email, "me@example.com")
  t.Equal(client.Description, "Foo")
}

func (t *testSuite) TestClientUpdate() {
  client, err := t.ApiClient.ClientUpdate(clientId, "you@example.com", "Bar")

  t.Equal(err, nil)

  t.Equal(client.Email, "you@example.com")
  t.Equal(client.Description, "Bar")
}
