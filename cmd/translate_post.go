package main

import (
	"github.com/dasuken/wizards-client/api/pkg/applications"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(applications.TranslatePost)
}

