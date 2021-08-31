package client

import (
	"github.com/stripe/stripe-go"
)

type PaymentClient interface {
	ProcessPayment(*stripe.ChargeParams) (string, error)
}
