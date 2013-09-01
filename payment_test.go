package paymill

import (
)

func (t *testSuite) TestCreateCreditCardPayment() {
  token := t.Config.Fixtures.Payment.Token
  c := NewApiClient(t.Config.Api.Token)

  p, err := c.CreatePayment(token, nil)

  t.Equal(err, nil)

  t.Not(t.Equal(p.Id, ""))
  t.Equal(p.PaymentType(), CreditCard)
}

func (t *testSuite) TestPaymentDetailsForNonExistingPayment() {
  id := "not_found"
  c := NewApiClient(t.Config.Api.Token)

  _, err := c.PaymentDetails(id)

  t.Equal("Payment not Found", err.Error())
}

func (t *testSuite) TestListPayments() {
  c := NewApiClient(t.Config.Api.Token)

  p, err := c.ListPayments("count", "card_type=visa")

  t.Equal(err, nil)

  t.Equal(len(p), 20)
}

func (t *testSuite) TestRemoveNonExistingPayment() {
  id := "pay_12345"
  c := NewApiClient(t.Config.Api.Token)

  ok, err := c.DeletePayment(id)
  t.Equal("Creditcard not found", err.Error())
  t.Equal(ok, false)
}
