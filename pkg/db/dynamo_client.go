package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"os"
)

/*
sessionの扱い

handlerとcontrollerの違い
*/

var DefaultDynamoClient dynamodbiface.DynamoDBAPI

func init() {
	defaultRegion := os.Getenv("AWS_REGION")
	if len(defaultRegion) == 0 {
		defaultRegion = "ap-northeast-1"
	}

	config := &aws.Config{
		Region: aws.String(defaultRegion),
	}

	env := os.Getenv("ENV")
	if env == "TEST" {
		config.Endpoint = aws.String("http://localhost:8000")
	}

	DefaultDynamoClient = NewClient(config)
}

func NewClient(config *aws.Config) dynamodbiface.DynamoDBAPI {
	return dynamodb.New(session.Must(session.NewSession(config)))
}