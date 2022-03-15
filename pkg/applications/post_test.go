package applications

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchPosts(t *testing.T) {
	testCases := []struct {
		name     string
		expected interface{}
		args     map[string]string
	}{
		{
			name:     "successfully fetch posts with limit",
			expected: 200,
			args:     map[string]string{
				"subreddit": "washingtonwizards",
				"limit": "1",
			},
		},
		{
			name:     "unsuccessfully fetch posts without subreddit",
			expected: 400,
			args:     map[string]string{
				"limit": "1",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := FetchPosts(buildParams(tc.args))

			assert.NoError(t, err)

			assert.Equal(t, tc.expected, res.StatusCode)
		})
	}
}


func TestFetchPostComments(t *testing.T) {
	testCases := []struct {
		name     string
		expected interface{}
		args     map[string]string
	}{
		{
			name:     "successfully fetchPostComments with exist postId",
			expected: 200,
			args:     map[string]string{
				"id": "t364q8",
			},
		},
		{
			name:     "unsuccessfully fetchPostComments with nonexist postId",
			expected: 400,
			args:     map[string]string{
				"id": "1",
			},
		},
		{
			name:     "unsuccessfully fetchPostComments without postId",
			expected: 400,
			args:     map[string]string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := FetchPostComments(buildParams(tc.args))

			assert.NoError(t, err)

			assert.Equal(t, tc.expected, res.StatusCode)
		})
	}
}

func buildParams(params map[string]string) events.APIGatewayProxyRequest {
	return events.APIGatewayProxyRequest{
		QueryStringParameters: params,
	}
}

