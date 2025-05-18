package main

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"buyallmemes.com/blog-api/src/domain/blog"
	"buyallmemes.com/blog-api/src/infrastructure/logging"
	"github.com/aws/aws-lambda-go/events"
	"github.com/mfenderov/konfig"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Initialize the logger for tests
	logger = logging.New(logging.DefaultConfig())

	// Load configuration
	if err := konfig.Load(); err != nil {
		logger.Error("Failed to load application properties", "error", err)
		os.Exit(1)
	}

	// Run tests
	os.Exit(m.Run())
}

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
