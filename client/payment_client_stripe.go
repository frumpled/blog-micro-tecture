package client

import (
	"app/model"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

type StripeClient struct{}

func NewPaymentClient(stripeKey string) PaymentClient {
	stripe.Key = stripeKey

	return StripeClient{}
}

func (s StripeClient) ProcessPayment(
	c model.Payment,
) (string, error) {
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(c.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Source:   &stripe.SourceParams{Token: stripe.String(c.StripeToken)},
	}
	params.AddMetadata("key", "value")

	ch, err := charge.New(params)

	if err != nil {
		return "", err
	}

	cp := stripe.CaptureParams{
		Amount: params.Amount,
	}

	ch, err = charge.Capture(ch.ID, &cp)

	return ch.ID, err
}
