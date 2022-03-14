package persistence

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type PostTable struct {
	dynamodbiface.DynamoDBAPI
	tableName string
}

func NewPostTable(client dynamodbiface.DynamoDBAPI, tableName string) (*PostTable, error) {
	return &PostTable{
		client,
		tableName,
	}, nil
}

type PostJAInfo struct {
	ID      string `json:"id"`
	BodyJA  string `json:"body_ja"`
	TitleJA string `json:"title_ja"`
}

func (t *PostTable) PutOne(post *PostJAInfo) error {
	av, err := dynamodbattribute.MarshalMap(post)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String(t.tableName),
	}

	_, err = t.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (t *PostTable) GetByID(id string) (*PostJAInfo, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id":{
				S: aws.String(id),
			},
		},
		TableName: aws.String(t.tableName),
	}

	result, err := t.GetItem(input)
	if err != nil {
		return nil, err
	}

	item := new(PostJAInfo)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}
/*
aws dynamodb create-table \
    --table-name Comments \
    --attribute-definitions \
        AttributeName=ID,AttributeType=S \
    --key-schema \
        AttributeName=ID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
    --region ap-northeast-1 \
    --endpoint-url http://localhost:8000
*/
