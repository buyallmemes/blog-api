package main

import (
	"buyallmemes.com/blog-api/src/blog/fetcher"
	"buyallmemes.com/blog-api/src/blog/fetcher/github"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mfenderov/konfig"
	"github.com/pkg/errors"
	"log"
	"os/exec"
)

func init() {
	ls("ls")
	ls("pwd")
	err := konfig.LoadConfiguration("resources/application.yaml")
	if err != nil {
		log.Fatal(errors.Wrap(err, "error loading application properties"))
	}
}
func ls(command string) {
	app := command
	cmd := exec.Command(app)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Print the output
	fmt.Println(string(stdout))
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, _ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	blog := getPosts(ctx)
	body, err := json.Marshal(blog)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error marshalling blog posts",
			StatusCode: 500,
		}, errors.Wrap(err, "error marshalling blog posts")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func getPosts(ctx context.Context) *fetcher.Blog {
	backend := github.New()

	blogFetcher := fetcher.BlogFetcher{
		BlogProvider: backend,
	}
	return blogFetcher.Fetch(ctx)
}
