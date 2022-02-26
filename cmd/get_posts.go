package main

import (
	"github.com/dasuken/wizards-client/api/pkg/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlers.FetchPosts)
}

