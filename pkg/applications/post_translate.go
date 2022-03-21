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

func TranslatePost(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var post domain.Post
	if err := json.Unmarshal([]byte(req.Body), &post); err != nil {
		log.Printf("%+v\n", err)
		return ResponseError(404, ErrorFailedToUnmarshal)
	}

	postTableName := os.Getenv("DYNAMO_TABLE_POST")
	if len(postTableName) == 0 {
		return ResponseError(http.StatusInternalServerError, "$DYNAMO_TABLE_POST is empty")
	}
	postTable, err := persistence.NewPostTable(persistence.DefaultDynamoClient, postTableName)
	if err != nil {
		log.Printf("%+v\n", err)
		return ResponseError(500, ErrorCouldNotConnectDB)
	}

	postInfo, err := postTable.GetByID(post.ID)
	if err != nil  {
		log.Printf("%+v\n", err)
		return ResponseError(http.StatusInternalServerError, ErrorFailedToFetch)
	}

	// if record not found, translate and put item
	//if err == dynamo.ErrNotFound || len(postInfo.ID) == 0 {
	if len(postInfo.ID) == 0 {
		//trs := translator.New(translator.DefaultAwsClient)
		trs_deepl := translator.New(translator.DefaultDeepLClient)

		// タイトルだけはdeeplで
		post.TitleJA, err = trs_deepl.Do(post.Title)
		if err != nil {
			log.Printf("%+v\n", err)
		}
		// to save cost
		if len(post.Body) < 1000 {
			post.BodyJA, err = trs_deepl.Do(post.Body)
			if err != nil {
				log.Printf("%+v\n", err)
			}
		}

		err = postTable.PutOne(&persistence.PostJAInfo{
			ID:      post.ID,
			BodyJA:  post.BodyJA,
			TitleJA: post.TitleJA,
		})
		if err != nil {
			log.Printf("%+v\n", err)
			return ResponseError(http.StatusInternalServerError, ErrorCouldNotPutItem)
		}

		return ResponseJson(http.StatusOK, post)
	}

	post.TitleJA = postInfo.TitleJA
	post.BodyJA = postInfo.BodyJA

	return ResponseJson(http.StatusOK, post)
}
