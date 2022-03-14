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

type RequestTranslatePost struct {
	ID    string `json:"id"`
	Body  string `json:"body"`
	Title string `json:"title"`
}

type ResponseTranslatePost struct {
	ID      string `json:"id"`
	BodyJA  string `json:"body_ja"`
	TitleJA string `json:"title_ja"`
}

func TranslatePost(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var request RequestTranslatePost
	if err := json.Unmarshal([]byte(req.Body), &request); err != nil {
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

	var response ResponseTranslatePost
	postInfo, err := postTable.GetByID(request.ID)
	if err != nil  {
		log.Printf("%+v\n", err)
		return ResponseError(http.StatusInternalServerError, ErrorFailedToFetch)
	}

	// if record not found, translate and put item
	//if err == dynamo.ErrNotFound || len(postInfo.ID) == 0 {
	if len(postInfo.ID) == 0 {
		//trs := translator.New(translator.DefaultAwsClient)
		trs_deepl := translator.New(translator.DefaultDeepLClient)

		response.ID = request.ID
		// タイトルだけはdeeplで
		response.TitleJA, err = trs_deepl.Do(request.Title)
		if err != nil {
			log.Printf("%+v\n", err)
		}
		// to save cost
		if len(request.Body) < 1000 {
			response.BodyJA, err = trs_deepl.Do(request.Body)
			if err != nil {
				log.Printf("%+v\n", err)
			}
		}

		err = postTable.PutOne(&persistence.PostJAInfo{
			ID:      response.ID,
			BodyJA:  response.BodyJA,
			TitleJA: response.TitleJA,
		})
		if err != nil {
			log.Printf("%+v\n", err)
			return ResponseError(http.StatusInternalServerError, ErrorCouldNotPutItem)
		}

		return ResponseJson(http.StatusOK, response)
	}

	response.ID = postInfo.ID
	response.TitleJA = postInfo.TitleJA
	response.BodyJA = postInfo.BodyJA

	return ResponseJson(http.StatusOK, response)
}
