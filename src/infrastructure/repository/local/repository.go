package local

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"buyallmemes.com/blog-api/src/domain/blog"
)

// LocalRepository implements the PostRepository interface using the local filesystem
type LocalRepository struct {
	markdownParser blog.MarkdownParser
	postsPath      string
}

// NewLocalRepository creates a new LocalRepository instance
func NewLocalRepository(markdownParser blog.MarkdownParser, postsPath string) *LocalRepository {
	return &LocalRepository{
		markdownParser: markdownParser,
		postsPath:      postsPath,
	}
}

// FetchPosts fetches all blog posts from the local filesystem
func (r *LocalRepository) FetchPosts(ctx context.Context) ([]blog.Post, error) {
	dir, err := os.ReadDir(r.postsPath)
	if err != nil {
		return nil, fmt.Errorf("error reading posts directory: %w", err)
	}

	var wg sync.WaitGroup
	posts := make([]blog.Post, 0)
	postsChan := make(chan blog.Post)
	errorsChan := make(chan error, len(dir))
	done := make(chan bool)

	// Start a goroutine to collect results
	go func() {
		for post := range postsChan {
			posts = append(posts, post)
		}
		done <- true
	}()

	// Create a buffered channel to act as a job queue
	jobQueue := make(chan os.DirEntry, len(dir))

	// Launch a fixed number of worker goroutines
	numWorkers := 5 // Adjust this number based on system capacity
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range jobQueue {
				post, err := r.fetchPost(file)
				if err != nil {
					errorsChan <- err
					continue
				}
				postsChan <- post
			}
		}()
	}

	// Enqueue jobs for workers
	for _, file := range dir {
		if !strings.HasSuffix(file.Name(), ".md") {
			continue
		}
		jobQueue <- file
	}
	close(jobQueue)

	// Wait for all goroutines to finish and close channels
	go func() {
		wg.Wait()
		close(postsChan)
	}()

	// Wait for either all posts to be collected or an error
	select {
	case <-done:
		return posts, nil
	case err := <-errorsChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// fetchPost fetches a single post from the local filesystem
func (r *LocalRepository) fetchPost(file os.DirEntry) (blog.Post, error) {
	content, err := r.getPostContent(file)
	if err != nil {
		return blog.Post{}, err
	}

	parsed, err := r.markdownParser.ParseMarkdown(content)
	if err != nil {
		return blog.Post{}, err
	}

	return blog.Post{
		Filename: file.Name(),
		Content:  parsed.Content,
		Date:     parsed.Date,
		Title:    parsed.Title,
		Anchor:   parsed.Anchor,
	}, nil
}

// getPostContent gets the content of a post from the local filesystem
func (r *LocalRepository) getPostContent(file os.DirEntry) (string, error) {
	filePath := filepath.Join(r.postsPath, file.Name())
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filePath, err)
	}
	return string(fileContent), nil
}
