package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

type DBModel struct {
	DBSession *dynamodb.DynamoDB
	TableName string
}

func NewDBModel(tableName string) DBModel {
	return DBModel{DBSession: DynamoSession(), TableName: tableName}
}

func (model DBModel) PutRecord(data interface{}) bool {
	dbData, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		log.Fatal("Failed to convert data", err)
		return false
	}

	_, err = model.DBSession.PutItem(&dynamodb.PutItemInput{
		Item:      dbData,
		TableName: aws.String(model.TableName),
	})

	if err != nil {
		log.Fatal("Failed to save data", err)
		return false
	}

	return true
}

func (model DBModel) ListRecords() *dynamodb.ScanOutput {
	params := &dynamodb.ScanInput{TableName: aws.String(model.TableName)}
	resp, err := model.DBSession.Scan(params)
	if err != nil {
		log.Fatal("Failed to reterive results", err)
	}
	return resp
}
