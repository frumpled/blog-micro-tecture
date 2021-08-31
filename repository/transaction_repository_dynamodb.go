package repository

import (
	repository_model "app/repository/model"
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
	transaction repository_model.Transaction,
) error {
	return t.ddbClient.save(transaction)
}
