package persistence

import (
	"fmt"
	"testing"
)

func testCommentTable() (*CommentTable, error) {
	return NewCommentTable(testClient(), "translate_comments_dev")
}

//
func TestPutComment(t *testing.T) {

	commentTable, err := testCommentTable()
	if err != nil {
		t.Fatal(err)
	}

	item := &CommentJAInfo{
		ID:     "comment3",
		PostID: "post3",
		BodyJA: "おはよう",
	}

	err = commentTable.PutOne(item)
	if err != nil {
		t.Fatal(err)
	}

	getItem, err := commentTable.GetByID(item.ID, item.PostID)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(getItem)
}

func TestGetComment(t *testing.T) {
	commentTable, err := testCommentTable()
	if err != nil {
		t.Fatal(err)
	}

	getItem, err := commentTable.GetByID("testid1", "post1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("item.")
	fmt.Println(getItem)
}
