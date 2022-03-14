package domain

import go_reddit "github.com/caarlos0/go-reddit/v3/reddit"

type Comment struct {
	go_reddit.Comment
	PostID  string `json:"post_id"`
	BodyJA  string `json:"body_ja"`
}
