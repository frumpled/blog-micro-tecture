package repository

import (
	stripe "github.com/stripe/stripe-go"
)

func NewTransactionRepository() TransactionRepository {
	return transactionRepositoryDynamoDB{
		ddbClient: newDDBClient(),
	}
}

type transactionRepositoryDynamoDB struct {
	ddbClient ddbClient
}

func (t transactionRepositoryDynamoDB) Save(
	transaction stripe.Transfer,
) error {
	return t.ddbClient.save(transaction)
}
