package model

type CreatePaymentRequest struct {
	Amount      int64  `json:"amount"`
	StripeToken string `json:"stripe_token"`
}
