package db

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"testing"
)

func testClient() dynamodbiface.DynamoDBAPI {
	return NewClient(&aws.Config{
		Region:aws.String("ap-northeast-1"),
		Endpoint:aws.String("http://localhost:8000"),
	})
}

func testPostTable() (*PostTable, error) {
	return NewPostTable(testClient(), "translate_posts_dev")
}
//
func TestPutPost(t *testing.T) {
	postTable, err := testPostTable()
	if err != nil {
		t.Fatal(err)
	}

	item := &PostJAInfo{
		ID:      "3",
		BodyJA:  "おはよう",
		TitleJA: "世界",
	}

	err = postTable.PutOne(item)
	if err != nil {
		t.Fatal(err)
	}

	getItem, err := postTable.GetByID(item.ID)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(getItem)
}

func TestGetPost(t *testing.T) {
	postTable, err := testPostTable()
	if err != nil {
		t.Fatal(err)
	}

	getItem, err := postTable.GetByID("1")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(getItem)
}