package utils

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

// return an AWS session ...
func GetAWSSession() *session.Session {
	sess, err := session.NewSession()
	if err != nil {
		// Handle Session creation error
		log.Fatal("Error in creating session")
	}

	return &sess
}

func DynamoSession() {
	dynamodb.New(GetAWSSession())
}
