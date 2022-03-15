package domain

type Post struct {
	ID      string `json:"id"`
	Body    string `json:"body"`
	Title   string `json:"title"`
	BodyJA  string `json:"body_ja"`
	TitleJA string `json:"title_ja"`
}
