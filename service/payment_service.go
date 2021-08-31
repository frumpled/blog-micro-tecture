package service

import (
	"app/client"
	controller_model "app/controller/model"
	"app/repository"
	repository_model "app/repository/model"

	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go"
)

// PaymentService provides core payment processing functions
type PaymentService interface {
	ProcessPayment(controller_model.ProcessPaymentRequest) (string, error)
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
	payment controller_model.ProcessPaymentRequest,
) (string, error) {
	stripeChargeParams := &stripe.ChargeParams{
		Amount:   stripe.Int64(payment.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Source:   &stripe.SourceParams{Token: stripe.String(payment.StripeToken)},
	}

	chargeID, err := p.paymentClient.ProcessPayment(stripeChargeParams)

	if err != nil {
		log.Println("Error processing payment: " + err.Error())
		return chargeID, err
	}

	transaction := createTransaction(payment.Amount, chargeID)
	err = p.transactionRepository.Save(transaction)

	return chargeID, err
}

func createTransaction(
	amount int64,
	vendorTransactionID string,
) repository_model.Transaction {
	id := uuid.NewString()
	createdAt := time.Now().Unix()

	partitionKey := fmt.Sprintf("%s%s%s", repository.TABLE_PREFIX_TRANSACTION, repository.TABLE_DELIMITER, id)
	sortKey := strconv.FormatInt(createdAt, 10)

	return repository_model.Transaction{
		PartitionKey:        partitionKey,
		SortKey:             sortKey,
		Amount:              amount,
		VendorTransactionID: vendorTransactionID,
	}
}
