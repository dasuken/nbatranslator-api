package reddit

import "testing"

const team = "washingtonwizards"

func Test_FetchPostList(t *testing.T) {
	posts, err := FetchPostList(team, Option{Limit: 3})
	if err != nil {
		t.Error(err)
	}

	if len(posts) != 3 {
		t.Error("posts length should be 3. but got", len(posts))
	}
}
