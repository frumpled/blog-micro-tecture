package client

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

type StripeClient struct{}

func NewPaymentClient(stripeKey string) PaymentClient {
	stripe.Key = stripeKey

	return StripeClient{}
}

func (s StripeClient) ProcessPayment(
	c *stripe.ChargeParams,
) (string, error) {

	ch, err := charge.New(c)

	if err != nil {
		return "", err
	}

	cp := stripe.CaptureParams{
		Amount: c.Amount,
	}

	ch, err = charge.Capture(ch.ID, &cp)

	return ch.ID, err
}
