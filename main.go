package main

import (
	"buyallmemes.com/blog-api/src/blog/fetcher"
	"buyallmemes.com/blog-api/src/blog/fetcher/github"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mfenderov/konfig"
	"github.com/pkg/errors"
	"log"
)

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

func init() {
	err := konfig.LoadConfiguration("resources/application.yaml")
	if err != nil {
		log.Fatal(errors.Wrap(err, "error loading application properties"))
	}
}

func main() {
	lambda.Start(handler)
}

func getPosts(ctx context.Context) *fetcher.Blog {
	backend := github.New()

	blogFetcher := fetcher.BlogFetcher{
		BlogProvider: backend,
	}
	return blogFetcher.Fetch(ctx)
}
