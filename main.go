package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"buyallmemes.com/blog-api/src/infrastructure/logging"
	"buyallmemes.com/blog-api/src/infrastructure/markdown"
	"buyallmemes.com/blog-api/src/infrastructure/repository/github"
	blogUsecase "buyallmemes.com/blog-api/src/usecase/blog"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mfenderov/konfig"
	"github.com/pkg/errors"
)

// Global logger instance
var logger *logging.Logger

// Application errors
var (
	ErrConfigLoading   = errors.New("configuration loading failed")
	ErrServiceCreation = errors.New("service creation failed")
	ErrFetchingPosts   = errors.New("error fetching blog posts")
	ErrMarshallingJSON = errors.New("error marshalling JSON")
)

// Configuration keys
const (
	GitHubOwnerKey = "github.owner"
	GitHubRepoKey  = "github.repo"
	GitHubPathKey  = "github.path"
	GitHubTokenKey = "github.token"
	DebugModeKey   = "debug.mode"
)

// Default values
const (
	DefaultGitHubOwner = "buyallmemes"
	DefaultGitHubRepo  = "blog-api"
	DefaultGitHubPath  = "posts"
	DefaultTimeout     = 30 * time.Second
)

func init() {
	// Initialize the logger
	logConfig := logging.DefaultConfig()
	if isDebugMode() {
		logConfig.Level = logging.DebugLevel
		logConfig.AddSource = true
	}
	logger = logging.New(logConfig)

	// Load configuration
	if err := konfig.Load(); err != nil {
		logger.Error("Failed to load application properties", "error", err)
		os.Exit(1)
	}
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()

	// Create the blog service
	blogService, err := createBlogService()
	if err != nil {
		logger.Error("Error creating blog service", "error", err)
		return createErrorResponse(http.StatusInternalServerError, "Internal server error"),
			errors.Wrap(err, "error creating blog service")
	}

	// Get all blog posts
	blogData, err := blogService.GetAllPosts(ctx)
	if err != nil {
		logger.Error("Error fetching blog posts", "error", err)
		return createErrorResponse(http.StatusInternalServerError, "Error fetching blog posts"),
			errors.Wrap(err, "error fetching blog posts")
	}

	// Marshal the blog data to JSON
	body, err := json.Marshal(blogData)
	if err != nil {
		logger.Error("Error marshalling blog posts", "error", err)
		return createErrorResponse(http.StatusInternalServerError, "Error processing blog posts"),
			errors.Wrap(err, "error marshalling blog posts")
	}

	// Return the successful response
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

// createBlogService creates and configures the blog service with its dependencies
func createBlogService() (blogUsecase.BlogService, error) {
	// Create the markdown parser
	markdownParser := markdown.NewGoldmarkParser()

	// Get GitHub configuration from environment variables
	config := github.NewConfig(
		getEnvWithDefault(GitHubOwnerKey, DefaultGitHubOwner),
		getEnvWithDefault(GitHubRepoKey, DefaultGitHubRepo),
		getEnvWithDefault(GitHubPathKey, DefaultGitHubPath),
		konfig.GetEnv(GitHubTokenKey),
	)

	// Create the repository
	repository, err := github.NewGitHubRepository(markdownParser, config)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrServiceCreation, err)
	}

	// Create and return the blog service
	return blogUsecase.NewBlogService(repository), nil
}

// createErrorResponse creates an API Gateway response for error cases
func createErrorResponse(statusCode int, message string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

// isDebugMode checks if debug mode is enabled
func isDebugMode() bool {
	debugMode := konfig.GetEnv(DebugModeKey)
	return debugMode == "true" || debugMode == "1"
}

// getEnvWithDefault gets an environment variable with a default value
func getEnvWithDefault(key, defaultValue string) string {
	value := konfig.GetEnv(key)
	if value == "" {
		// Log that we're using a default value if debug mode is enabled
		if isDebugMode() {
			logger.Debug("Using default value", "key", key, "value", defaultValue)
		}
		return defaultValue
	}
	return value
}
