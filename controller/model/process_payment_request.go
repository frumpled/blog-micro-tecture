package model

type ProcessPaymentRequest struct {
	Amount      int64  `json:"amount"`
	StripeToken string `json:"stripe_token"`
}
