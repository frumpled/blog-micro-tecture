package model

type Payment struct {
	Amount      int64
	StripeToken string
	Description string
}
