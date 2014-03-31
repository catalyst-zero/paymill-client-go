package paymill

import ()

func (t *testSuite) TestCreateCreditCardPayment() {
	token := t.Config.Fixtures.Payment.Token

	p, err := t.ApiClient.CreatePayment(token, nil)

	t.Equal(err, nil)

	t.Not(t.Equal(p.Id, ""))
}

func (t *testSuite) TestPaymentDetailsForNonExistingPayment() {
	id := "not_found"

	_, err := t.ApiClient.PaymentDetails(id)

	t.Equal("Payment not Found", err.Error())
}

func (t *testSuite) TestListPayments() {
	p, err := t.ApiClient.ListPayments("count", "card_type=visa")

	t.Equal(err, nil)

	t.Equal(len(p), 20)
}

func (t *testSuite) TestRemoveNonExistingPayment() {
	id := "pay_12345"

	ok, err := t.ApiClient.DeletePayment(id)
	t.Equal("Creditcard not found", err.Error())
	t.Equal(ok, false)
}
