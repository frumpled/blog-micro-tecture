package service

import (
	"app/client"
	"app/model"
	"app/repository"

	"log"
	"time"

	"github.com/google/uuid"
)

// PaymentService provides core payment processing functions
type PaymentService interface {
	ProcessPayment(model.Payment) (string, error)
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
	payment model.Payment,
) (string, error) {
	chargeID, err := p.paymentClient.ProcessPayment(payment)

	if err != nil {
		log.Println("Error processing payment: " + err.Error())
		return chargeID, err
	}

	transaction := createTransaction(payment.Amount, chargeID, payment.Description)
	err = p.transactionRepository.Save(transaction)

	return chargeID, err
}

func createTransaction(
	amount int64,
	vendorTransactionID string,
	description string,
) model.Transaction {
	return model.Transaction{
		ID:                  uuid.NewString(),
		CreatedAt:           time.Now().Unix(),
		Amount:              amount,
		VendorTransactionID: vendorTransactionID,
		Description:         description,
	}
}
