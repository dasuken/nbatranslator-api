package post

import (
	"context"
	"errors"
	"fmt"
	"github.com/caarlos0/go-reddit/v3/reddit"
)

var(
	ErrorFailedToFetchPosts = "failed to fetch posts"
	ErrorSubredditNotExist = "subreddit does not exist"
)

const defaultPostsLimit = 10

type Option struct {
	Limit int
}

func FetchAll(subreddit string, option Option) ([]*reddit.Post, error) {
	if subreddit == "" {
		return nil, errors.New(ErrorSubredditNotExist)
	}

	if option.Limit == 0 {
		option.Limit = defaultPostsLimit
	}

	posts, _, err := reddit.DefaultClient().Subreddit.NewPosts(context.Background(), subreddit, &reddit.ListOptions{
		Limit: option.Limit,
	})

	if err != nil {
		return nil, fmt.Errorf("error message: %v, received values: %v", err, posts)
	}

	return posts, nil
}

func FetchOne(postID string) (*reddit.PostAndComments, error) {
	postComments, _, err := reddit.DefaultClient().Post.Get(context.Background(), postID)
	if err != nil {
		return nil, err
	}

	comments := postComments.Comments
	if len(comments) > 10 {
		postComments.Comments = comments[:10]
	}

	return postComments, nil
}
