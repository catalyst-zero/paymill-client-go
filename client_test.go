package paymill

func (t *testSuite) TestClient() {
  client, err := t.ApiClient.CreateClient("me@example.com", "Foo")

  t.Equal(err, nil)

  t.Not(t.Equal(client.Id, ""))
  t.Equal(client.Email, "me@example.com")
  t.Equal(client.Description, "Foo")

  t.ApiClient.RemoveClient(client.Id)
}

func (t *testSuite) TestClientDetails() {
  c, err := t.ApiClient.CreateClient("me@example.com", "Foo")

  client, err := t.ApiClient.ClientDetails(c.Id)

  t.Equal(err, nil)

  t.Equal(client.Email, "me@example.com")
  t.Equal(client.Description, "Foo")

  t.ApiClient.RemoveClient(client.Id)
}

func (t *testSuite) TestClientUpdate() {
  c, err := t.ApiClient.CreateClient("me@example.com", "Foo")

  client, err := t.ApiClient.ClientUpdate(c.Id, "you@example.com", "Bar")

  t.Equal(err, nil)

  t.Equal(client.Email, "you@example.com")
  t.Equal(client.Description, "Bar")

  t.ApiClient.RemoveClient(client.Id)
}

func (t *testSuite) TestRemoveClient() {
  client, err := t.ApiClient.CreateClient("me@example.com", "Foo")
  ok, err := t.ApiClient.RemoveClient(client.Id)

  t.Equal(err, nil)
  t.Equal(ok, true)
}
