package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dasuken/wizards-client/api/pkg/applications"
)

func main() {
	lambda.Start(applications.FetchPostComments)
}
