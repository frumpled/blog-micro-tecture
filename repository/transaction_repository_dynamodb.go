package repository

import (
	"app/model"
	repository_model "app/repository/model"

	"fmt"
	"strconv"
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
	transaction model.Transaction,
) error {
	partitionKey := fmt.Sprintf("%s%s%s", TABLE_PREFIX_TRANSACTION, TABLE_DELIMITER, transaction.ID)
	sortKey := strconv.FormatInt(transaction.CreatedAt, 10)

	txnData := repository_model.Transaction{
		PartitionKey:        partitionKey,
		SortKey:             sortKey,
		VendorTransactionID: transaction.VendorTransactionID,
		Amount:              transaction.Amount,
		Description:         transaction.Description,
	}

	return t.ddbClient.save(txnData)
}
