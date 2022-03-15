package applications

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dasuken/wizards-client/api/pkg/domain"
	"github.com/dasuken/wizards-client/api/pkg/infra/persistence"
	"github.com/dasuken/wizards-client/api/pkg/infra/translator"
	"log"
	"net/http"
	"os"
)

func TranslateComment(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var comment domain.Comment
	if err := json.Unmarshal([]byte(req.Body), &comment); err != nil {
		log.Printf("%+v\n", err)
		return ResponseError(404, ErrorFailedToUnmarshal)
	}

	if len(comment.ID) == 0 {
		return ResponseError(http.StatusBadRequest, ErrorCouldNotEmptyID)
	}

	if len(comment.PostID) == 0 {
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
	commentInfo, err := commentTable.GetByID(comment.ID, comment.PostID)
	log.Printf("GetByID(%s), commentInfo: %v, err: %v", comment.ID, commentInfo, err)

	// if record not found, translate and put item
	if len(commentInfo.ID) == 0 {
		trs := translator.New(translator.DefaultAwsClient)
		if len(comment.Body) < 1000 {
			comment.BodyJA, err = trs.Do(comment.Body)
			if err != nil {
				log.Printf("%+v\n", err)
			}
		}

		err = commentTable.PutOne(&persistence.CommentJAInfo{
			ID:      comment.ID,
			PostID:  comment.PostID,
			BodyJA:  comment.BodyJA,
		})
		if err != nil {
			log.Printf("%+v\n", err)
			return ResponseError(http.StatusInternalServerError, ErrorCouldNotPutItem)
		}

		return ResponseJson(http.StatusOK, comment)
	}

	// else error
	if err != nil {
		log.Printf("%+v\n", err)
		return ResponseError(500, ErrorFailedToFetch)
	}

	comment.BodyJA = commentInfo.BodyJA

	return ResponseJson(http.StatusOK, comment)
}
