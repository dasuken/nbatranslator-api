.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/translate_post cmd/translate_post.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/translate_comment cmd/translate_comment.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get_posts cmd/get_posts.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get_post_comments cmd/get_post_comments.go

clean:
	rm -rf ./bin

deploy-dev: clean build
	sls deploy --verbose --stage="dev"

deploy-prd: clean build
	sls deploy --verbose --stage="prd"

dynamo-local:
	sls dynamodb start
