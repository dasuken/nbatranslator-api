package post

import (
	"testing"
)

func TestFetchAll(t *testing.T) {
	type arg struct {
		subreddit string
		option    Option
	}

	testCases := []struct {
		name     string
		arg      arg
	}{
		{
			name:     "should return posts",
			arg: arg{
				"washingtonwizards",
				Option{Limit: 3},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := FetchAll(tc.arg.subreddit, tc.arg.option)
			if err != nil {
				t.Error(err)
			}

			if len(out) != tc.arg.option.Limit {
				t.Errorf("unexpected response length: %d, expected: %d", len(out), tc.arg.option.Limit)
			}
		})
	}
}
