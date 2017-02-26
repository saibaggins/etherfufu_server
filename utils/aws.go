package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

// return an AWS session ...
func GetAWSSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		// Handle Session creation error
		log.Fatal("Error in creating session")
	}

	return sess
}

func DynamoSession() *dynamodb.DynamoDB {
	return dynamodb.New(GetAWSSession())
}

func S3Service() *s3.S3 {
	return s3.New(GetAWSSession())
}
