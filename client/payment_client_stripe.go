package client

import (
	"app/model"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type StripeClient struct{}

func NewPaymentClient(stripeKey string) PaymentClient {
	stripe.Key = stripeKey

	return StripeClient{}
}

func (s StripeClient) ProcessPayment(
	c model.Payment,
) (string, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(c.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
		Source:  stripe.String(c.StripeToken),
		Confirm: stripe.Bool(true),
	}
	params.AddMetadata("key", "value")
	pi, err := paymentintent.New(params)
	if err != nil {
		return "", err
	}

	return pi.ID, err
}
