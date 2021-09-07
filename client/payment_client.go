package client

import (
	"github.com/stripe/stripe-go"
)

type PaymentClient interface {
	ProcessPayment(*stripe.PaymentIntentParams) (string, error)
}
