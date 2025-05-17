package main

import (
	"context"
	"encoding/json"
	"testing"

	"buyallmemes.com/blog-api/src/domain/blog"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func Test_handler(t *testing.T) {
	response, err := handler(context.Background(), events.APIGatewayProxyRequest{})
	assert.Equal(t, 200, response.StatusCode)
	assert.NoError(t, err)
	body := response.Body

	blogData := blog.Blog{}
	unmarshallingError := json.Unmarshal([]byte(body), &blogData)
	assert.NoError(t, unmarshallingError)
	posts := blogData.Posts
	assert.True(t, len(posts) > 0)
	assert.Equal(t, "Hello, World!", posts[len(posts)-1].Title)
}
