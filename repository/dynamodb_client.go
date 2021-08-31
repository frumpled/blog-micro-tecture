package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dynamoDBClient *dynamodb.DynamoDB

func init() {
	session, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	dynamoDBClient = dynamodb.New(session)
}

type ddbClient struct {
	ddb *dynamodb.DynamoDB
}

func newDDBClient() ddbClient {
	return ddbClient{
		ddb: dynamoDBClient,
	}
}

func (d ddbClient) save(input interface{}) error {
	avs, err := dynamodbattribute.MarshalMap(
		input,
	)
	if err != nil {
		return err
	}

	_, err = d.ddb.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item:      avs,
	})

	return err
}
