package domain

import (
	go_reddit "github.com/caarlos0/go-reddit/v3/reddit"
)

type Post struct {
	*go_reddit.Post
	BodyJA  string `json:"body_ja"`
	TitleJA string `json:"title_ja"`
}
