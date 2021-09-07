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
	p *stripe.PaymentIntentParams,
) (string, error) {
	pi, err := paymentintent.New(p)
	if err != nil {
		return "", err
	}

	return pi.ID, err
}
