package applications

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dasuken/wizards-client/api/pkg/infra/persistence"
	"github.com/dasuken/wizards-client/api/pkg/infra/translator"
	"log"
	"net/http"
	"os"
)

type RequestTranslateComment struct {
	ID    string `json:"id"`
	PostID    string `json:"post_id"`
	Body  string `json:"body"`
}

type ResponseTranslateComment struct {
	ID      string `json:"id"`
	PostID  string `json:"post_id"`
	BodyJA  string `json:"body_ja"`
}

func TranslateComment(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var request *RequestTranslateComment
	if err := json.Unmarshal([]byte(req.Body), &request); err != nil {
		log.Printf("%+v\n", err)
		return ResponseError(404, ErrorFailedToUnmarshal)
	}

	if len(request.ID) == 0 {
		return ResponseError(http.StatusBadRequest, ErrorCouldNotEmptyID)
	}

	if len(request.PostID) == 0 {
		return ResponseError(http.StatusBadRequest, ErrorCouldNotEmptyPostID)
	}

	commentTableName := os.Getenv("DYNAMO_TABLE_COMMENT")
	if len(commentTableName) == 0 {
		return ResponseError(http.StatusInternalServerError, "$DYNAMO_TABLE_COMMENT is empty")
	}
	commentTable, err := persistence.NewCommentTable(persistence.DefaultDynamoClient, commentTableName)
	if err != nil {
		log.Printf("%+v\n", err)
		return ResponseError(500, ErrorCouldNotConnectDB)
	}

	var response ResponseTranslateComment
	commentInfo, err := commentTable.GetByID(request.ID, request.PostID)
	log.Printf("GetByID(%s), commentInfo: %v, err: %v", request.ID, commentInfo, err)

	// if record not found, translate and put item
	if len(commentInfo.ID) == 0 {
		trs := translator.New(translator.DefaultAwsClient)
		response.ID = request.ID
		response.PostID = request.PostID
		if len(request.Body) < 1000 {
			response.BodyJA, err = trs.Do(request.Body)
			if err != nil {
				log.Printf("%+v\n", err)
			}
		}

		err = commentTable.PutOne(&persistence.CommentJAInfo{
			ID:      response.ID,
			PostID:  response.PostID,
			BodyJA:  response.BodyJA,
		})
		if err != nil {
			log.Printf("%+v\n", err)
			return ResponseError(http.StatusInternalServerError, ErrorCouldNotPutItem)
		}

		return ResponseJson(http.StatusOK, response)
	}

	// else error
	if err != nil {
		log.Printf("%+v\n", err)
		return ResponseError(500, ErrorFailedToFetch)
	}

	response.ID = commentInfo.ID
	response.PostID = commentInfo.PostID
	response.BodyJA = commentInfo.BodyJA

	return ResponseJson(http.StatusOK, response)
}
