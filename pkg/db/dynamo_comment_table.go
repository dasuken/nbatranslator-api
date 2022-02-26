package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type CommentTable struct {
	dynamodbiface.DynamoDBAPI
	tableName string
}

func NewCommentTable(client dynamodbiface.DynamoDBAPI, tableName string) (*CommentTable, error) {
	return &CommentTable{
		client,
		tableName,
	}, nil
}

type CommentJAInfo struct {
	ID      string `json:"id"`
	PostID  string `json:"post_id"`
	BodyJA  string `json:"body_ja"`
}

func (t *CommentTable) GetByID(id, post_id string) (*CommentJAInfo, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id":{
				S: aws.String(id),
			},
			"post_id": {
				S: aws.String(post_id),
			},
		},
		TableName: aws.String(t.tableName),
	}

	result, err := t.GetItem(input)
	if err != nil {
		return nil, err
	}

	item := new(CommentJAInfo)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (t *CommentTable) PutOne(comment *CommentJAInfo) error {
	av, err := dynamodbattribute.MarshalMap(comment)
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