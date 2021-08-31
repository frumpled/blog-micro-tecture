package repository

import (
	"github.com/stripe/stripe-go"
)

type TransactionRepository interface {
	Save(stripe.Transfer) error
}
