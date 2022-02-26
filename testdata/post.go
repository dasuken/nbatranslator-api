package testdata

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/caarlos0/go-reddit/v3/reddit"
	"io"
	"os"
)

func loadPosts(subredditName string) ([]byte, error) {
	f, err := os.Open(fmt.Sprintf("testdata/%s_mock.json", subredditName))
	if err != nil {
		return nil, err
	}
	b, _ := io.ReadAll(f)
	return b, nil
}

func GetMockPosts(subredditName string) ([]reddit.Post, error) {
	var posts []reddit.Post
	b, err := loadPosts(subredditName)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func FetchPosts(subredditName string) error {
	posts, _, err := reddit.DefaultClient().Subreddit.HotPosts(context.Background(), subredditName, &reddit.ListOptions{
			Limit: 10,
		},
	)
	if err != nil {
		return err
	}

	f, _ := os.Create(fmt.Sprintf("testdata/%s_mock.json", subredditName))
	b, err := json.Marshal(posts)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	f.Write(b)
	return nil
}