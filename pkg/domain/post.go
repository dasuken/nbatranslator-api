package domain

import (
	go_reddit "github.com/caarlos0/go-reddit/v3/reddit"
)

type Post struct {
	*go_reddit.Post
	body_ja string
	title_ja string
}

