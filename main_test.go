package main

import (
	"buyallmemes.com/blog-api/src/blog/fetcher"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getPosts(t *testing.T) {

	response, err := handler(context.Background(), events.APIGatewayProxyRequest{})
	assert.Equal(t, 200, response.StatusCode)
	assert.NoError(t, err)
	body := response.Body

	blog := fetcher.Blog{}
	unmarshallingError := json.Unmarshal([]byte(body), &blog)
	assert.NoError(t, unmarshallingError)
	posts := blog.Posts
	assert.True(t, len(posts) > 0)
	assert.Equal(t, "Hello, World!", posts[len(posts)-1].Title)
}
