package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/dasuken/wizards-client/api/pkg/post"
	"github.com/dasuken/wizards-client/api/pkg/utils"
	"net/http"
)

var (
	ErrorInvalidLimit = "invalid limit"
	ErrorInvalidPostID = "invalid postID"
	ErrorInvalidSubreddit = "invalid subreddit"
)

// fetch and add kind property
func FetchPosts(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	subreddit := req.QueryStringParameters["subreddit"]
	if len(subreddit) == 0 {
		return ResponseError(http.StatusBadRequest, ErrorInvalidSubreddit)
	}

	limit := req.QueryStringParameters["limit"]
	limitNum, err := utils.ParseStr(limit)
	if err != nil {
		return ResponseError(http.StatusBadRequest, ErrorInvalidLimit)
	}

	posts, err := post.FetchAll(subreddit, post.Option{Limit: limitNum})
	if err != nil {
		return ResponseError(http.StatusBadRequest, err.Error())
	}

	return ResponseJson(http.StatusOK, posts)
}

func FetchPostComments(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	postID := req.QueryStringParameters["id"]
	if len(postID) == 0 {
		return ResponseError(http.StatusBadRequest, ErrorInvalidPostID)
	}

	postComments, err := post.FetchOne(postID)
	if err != nil {
		return ResponseError(http.StatusBadRequest, err.Error())
	}

	return ResponseJson(http.StatusOK, postComments)
}

