package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func commonHeaders() map[string]string {
	return map[string]string {
		"Content-Type": "application/json",
		"Access-Control-Allow-Origin": "*",
	}
}

func ResponseJson(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	b, _ := json.Marshal(body)

	resp := &events.APIGatewayProxyResponse{
		Headers: commonHeaders(),
		StatusCode: status,
		Body: string(b),
	}

	return resp, nil
}

func ResponseError(status int, message string) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		Headers:    commonHeaders(),
		StatusCode: status,
		Body:       fmt.Sprintf(`{"message":"%s"}`, message),
	}, nil
}
