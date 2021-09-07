package service

import (
	"app/client"
	"app/repository"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go"
)

// PaymentService provides core payment processing functions
type PaymentService interface {
	ProcessPayment(
		int64,
		string,
	) (string, error)
}

// NewPaymentService provides a service to interact with a payment processor backed by a PaymentClient, abstracting a 3rd party vendor
func NewPaymentService(
	paymentClient client.PaymentClient,
	transactionRepository repository.TransactionRepository,
) PaymentService {
	return paymentServiceStripe{
		paymentClient:         paymentClient,
		transactionRepository: transactionRepository,
	}
}

type paymentServiceStripe struct {
	paymentClient         client.PaymentClient
	transactionRepository repository.TransactionRepository
}

func (p paymentServiceStripe) ProcessPayment(
	amount int64,
	stripeToken string,
) (string, error) {
	paymentIntentID, err := p.paymentClient.ProcessPayment(
		amount,
		stripeToken,
	)

	if err != nil {
		log.Println("Error processing payment: " + err.Error())
		return paymentIntentID, err
	}

	transaction := createTransaction(amount, paymentIntentID)
	err = p.transactionRepository.Save(transaction)

	return paymentIntentID, err
}

func createTransaction(
	amount int64,
	vendorTransactionID string,
) stripe.Transfer {
	return stripe.Transfer{
		ID:      uuid.NewString(),
		Amount:  amount,
		Created: time.Now().Unix(),
		Metadata: map[string]string{
			"vendor_transaction_id": vendorTransactionID,
		},
	}
}
