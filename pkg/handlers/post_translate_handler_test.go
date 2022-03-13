package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TranslatePost(t *testing.T) {
	// req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error
	body := map[string]interface{}{
		"id": "testid1",
		"body": "hello",
		"title": "world",
	}
	bodyStr, err := json.Marshal(body)
	assert.NoError(t, err)

	res, err := TranslatePost(events.APIGatewayProxyRequest{
		Body: string(bodyStr),
	})

	// レスポンスコードをチェック
	assert.Equal(t, 200, res.StatusCode)

	// JSONからmap型に変換
	var resBody map[string]interface{}
	err = json.Unmarshal([]byte(res.Body), &resBody)
	assert.NoError(t, err)

	fmt.Println(res.Body)
}