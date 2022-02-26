package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dasuken/wizards-client/api/pkg/handlers"
)

func main() {
	lambda.Start(handlers.FetchPostComments)
}
