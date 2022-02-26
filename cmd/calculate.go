package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/go-reddit/v3/reddit"
	"log"
)

const subredditName = "nba"

func main() {
	//posts, err := post.FetchAll(subredditName, post.Option{Limit: 25})
	//if err != nil {
	//	log.Fatal()
	//}
	//
	//calculatePostTitleAndBodyLength(posts)
	calculateComment("sk0rp2")
}

func calculatePostTitleAndBodyLength(posts []*reddit.Post) {
	sumT := 0
	sumB := 0
	for _, post := range posts {
		sumT += len(post.Title)

		if len(post.Body) < 1000 {
			sumB += len(post.Body)
		}
	}

	fmt.Println("average title length!: ", sumT / len(posts))
	fmt.Println("average body length!: ", sumB / len(posts))
}

var cnt int

func calculateComment(postId string) {
	postComments, _, err := reddit.DefaultClient().Post.Get(context.Background(), postId)
	if err != nil {
		log.Fatal(err)
	}
	comments := postComments.Comments

	sum := 0
	for _, comment := range comments {
		fmt.Println(comment)
		fmt.Println(comment.HasMore())
		fmt.Println()
		sum += Recursive(comment)
	}
	fmt.Println(sum)
	fmt.Println(cnt)
}

func Recursive(comment *reddit.Comment) int {
	cnt ++
	if comment.Replies.Comments == nil {
		return len(comment.Body)
	}

	sum := 0
	for _, r := range comment.Replies.Comments {
		sum += Recursive(r)
	}

	return sum
}

