package domain

type Comment struct {
	ID     string `json:"id"`
	PostID string `json:"post_id"`
	Body   string `json:"body"`
	BodyJA string `json:"body_ja"`
}
