package domain

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

type Post struct {
	*reddit.Post
	body_ja string
	title_ja string
}

func GetPostList(subreddit string, option Option) ([]*Post, error) {
	if subreddit == "" {
		return nil, errors.New(ErrorSubredditNotExist)
	}

	if option.Limit == 0 {
		option.Limit = defaultPostsLimit
	}

	rawPosts, _, err := reddit.DefaultClient().Subreddit.NewPosts(context.Background(), subreddit, &reddit.ListOptions{
		Limit: option.Limit,
	})
	if err != nil {
		return nil, fmt.Errorf("error message: %v, received values: %v", err, rawPosts)
	}

	posts := make([]*Post, len(rawPosts))
	for i, post := range rawPosts {
		posts[i] = &Post{
			Post: post,
		}
	}

	return posts, nil
}

func GetPostByID(postID string) (*reddit.PostAndComments, error) {
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
