package client

type PaymentClient interface {
	ProcessPayment(int64, string) (string, error)
}
