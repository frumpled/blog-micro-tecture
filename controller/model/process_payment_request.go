package model

type ProcessPaymentRequest struct {
	Amount          int64  `json:"amount"`
	FundingSourceID string `json:"funding_source_id"`
	StripeToken     string `json:"stripe_token"`
	Description     string `json:"description"`
}
