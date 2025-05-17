package github

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"buyallmemes.com/blog-api/src/domain/blog"
	"github.com/google/go-github/v70/github"
)

// Common errors
var (
	ErrInvalidConfig    = errors.New("invalid repository configuration")
	ErrGitHubAPIFailure = errors.New("GitHub API failure")
	ErrContentDecoding  = errors.New("content decoding failure")
	ErrMarkdownParsing  = errors.New("markdown parsing failure")
	ErrContextCancelled = errors.New("context cancelled")
)

// GitHubRepository implements the PostRepository interface using the GitHub API
type GitHubRepository struct {
	client         *github.Client
	markdownParser blog.MarkdownParser
	config         *Config
}

// NewGitHubRepository creates a new GitHubRepository instance
func NewGitHubRepository(markdownParser blog.MarkdownParser, config *Config) (*GitHubRepository, error) {
	if markdownParser == nil {
		return nil, fmt.Errorf("%w: markdown parser is required", ErrInvalidConfig)
	}

	if config == nil {
		return nil, fmt.Errorf("%w: config is required", ErrInvalidConfig)
	}

	if config.Owner == "" || config.Repo == "" || config.Path == "" {
		return nil, fmt.Errorf("%w: owner, repo, and path are required", ErrInvalidConfig)
	}

	client := github.NewClient(nil)
	if config.Token != "" {
		client = client.WithAuthToken(config.Token)
	}

	return &GitHubRepository{
		client:         client,
		markdownParser: markdownParser,
		config:         config,
	}, nil
}

// FetchPosts fetches all blog posts from GitHub
func (r *GitHubRepository) FetchPosts(ctx context.Context) ([]blog.Post, error) {
	// Add timeout to context if not already set
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	directoryContent, err := r.getDirectoryContent(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get directory content: %w", err)
	}

	// Filter markdown files
	var mdFiles []*github.RepositoryContent
	for _, file := range directoryContent {
		if file != nil && file.Name != nil && strings.HasSuffix(*file.Name, ".md") {
			mdFiles = append(mdFiles, file)
		}
	}

	if len(mdFiles) == 0 {
		return []blog.Post{}, nil
	}

	// Use errgroup for better error handling
	var (
		mu    sync.Mutex
		posts []blog.Post
	)

	// Process files concurrently with a worker pool pattern
	errChan := make(chan error, 1)
	resultChan := make(chan blog.Post)

	// Start workers (limit concurrency to 5)
	const maxWorkers = 5
	workChan := make(chan *github.RepositoryContent, len(mdFiles))

	// Start a fixed number of workers
	var wg sync.WaitGroup
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range workChan {
				select {
				case <-ctx.Done():
					// Context cancelled, stop processing
					errChan <- fmt.Errorf("%w: %v", ErrContextCancelled, ctx.Err())
					return
				default:
					post, err := r.fetchPost(ctx, file)
					if err != nil {
						errChan <- err
						return
					}
					resultChan <- post
				}
			}
		}()
	}

	// Close work channel when all files are queued
	go func() {
		for _, file := range mdFiles {
			workChan <- file
		}
		close(workChan)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Process results and handle errors
	for {
		select {
		case err := <-errChan:
			// Cancel context to stop all workers
			cancel()
			return nil, err
		case post, ok := <-resultChan:
			if !ok {
				// All results processed
				return posts, nil
			}
			mu.Lock()
			posts = append(posts, post)
			mu.Unlock()
		case <-ctx.Done():
			return nil, fmt.Errorf("%w: %v", ErrContextCancelled, ctx.Err())
		}
	}
}

// fetchPost fetches a single post from GitHub
func (r *GitHubRepository) fetchPost(ctx context.Context, file *github.RepositoryContent) (blog.Post, error) {
	if file == nil || file.Name == nil {
		return blog.Post{}, fmt.Errorf("%w: invalid file", ErrInvalidConfig)
	}

	content, err := r.getPostContent(ctx, file.Name)
	if err != nil {
		return blog.Post{}, fmt.Errorf("failed to get post content: %w", err)
	}

	if content.Content == nil {
		return blog.Post{}, fmt.Errorf("%w: content is nil", ErrGitHubAPIFailure)
	}

	decoded, err := base64.StdEncoding.DecodeString(*content.Content)
	if err != nil {
		return blog.Post{}, fmt.Errorf("%w: %v", ErrContentDecoding, err)
	}

	parsed, err := r.markdownParser.ParseMarkdown(string(decoded))
	if err != nil {
		return blog.Post{}, fmt.Errorf("%w: %v", ErrMarkdownParsing, err)
	}

	return blog.Post{
		Filename: *file.Name,
		Content:  parsed.Content,
		Date:     parsed.Date,
		Title:    parsed.Title,
		Anchor:   parsed.Anchor,
	}, nil
}

// getDirectoryContent gets the content of a directory from GitHub
func (r *GitHubRepository) getDirectoryContent(ctx context.Context) ([]*github.RepositoryContent, error) {
	_, directoryContent, resp, err := r.client.Repositories.GetContents(
		ctx,
		r.config.Owner,
		r.config.Repo,
		r.config.Path,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrGitHubAPIFailure, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: non 200 response code: %v", ErrGitHubAPIFailure, resp.StatusCode)
	}

	return directoryContent, nil
}

// getPostContent gets the content of a post from GitHub
func (r *GitHubRepository) getPostContent(ctx context.Context, filename *string) (*github.RepositoryContent, error) {
	if filename == nil {
		return nil, fmt.Errorf("%w: filename is nil", ErrInvalidConfig)
	}

	content, _, resp, err := r.client.Repositories.GetContents(
		ctx,
		r.config.Owner,
		r.config.Repo,
		fmt.Sprintf("%s/%s", r.config.Path, *filename),
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrGitHubAPIFailure, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: non 200 response code: %v", ErrGitHubAPIFailure, resp.StatusCode)
	}

	return content, nil
}
