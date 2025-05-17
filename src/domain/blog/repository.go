package blog

import (
	"context"
)

// PostRepository defines the interface for fetching blog posts
type PostRepository interface {
	// FetchPosts fetches all blog posts
	FetchPosts(ctx context.Context) ([]Post, error)
}
