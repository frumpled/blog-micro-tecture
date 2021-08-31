package model

type Transaction struct {
	PartitionKey        string `dynamodbav:"pk"`
	SortKey             string `dynamodbav:"sk"`
	VendorTransactionID string `dynamodbav:"vid"`
	Amount              int64  `dynamodbav:"a"`
}
