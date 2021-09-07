package client

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type StripeClient struct{}

func NewPaymentClient(stripeKey string) PaymentClient {
	stripe.Key = stripeKey

	return StripeClient{}
}

func (s StripeClient) ProcessPayment(
	amount int64,
	stripeToken string,
) (string, error) {
	stripeParams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
		Source:  stripe.String(stripeToken),
		Confirm: stripe.Bool(true),
	}
	pi, err := paymentintent.New(stripeParams)
	if err != nil {
		return "", err
	}

	return pi.ID, err
}
