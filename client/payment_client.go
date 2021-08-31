package client

import "app/model"

type PaymentClient interface {
	ProcessPayment(model.Payment) (string, error)
}
