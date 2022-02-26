package testdata

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/caarlos0/go-reddit/v3/reddit"
	"github.com/dasuken/wizards-client/api/pkg/post"
	"io"
	"log"
	"os"
)

func loadComments(subredditName string) ([]byte, error) {
	f, err := os.Open(fmt.Sprintf("testdata/comment/%s_mock.json", subredditName))
	if err != nil {
		return nil, err
	}
	b, _ := io.ReadAll(f)
	return b, nil
}

func GetMockComments(subredditName string) ([]reddit.Post, error) {
	var posts []reddit.Post
	b, err := loadComments(subredditName)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func FetchComments(subredditName string) error {
	posts, err := post.FetchAll(subredditName, post.Option{Limit: 1})
	if err != nil {
		log.Fatal()
	}

	postComments, _, err := reddit.DefaultClient().Post.Get(context.Background(), posts[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	post := postComments.Post
	comments := postComments.Comments

	fmt.Println("post: \n", post)

	fmt.Println("comment: ")
	f, _ :=os.Create("res.json")
	b, _ :=json.Marshal(comments)
	f.Write(b)

	return nil
}
