package blog

import (
	"context"
	"sort"

	"buyallmemes.com/blog-api/src/domain/blog"
)

// BlogService defines the interface for blog-related use cases
type BlogService interface {
	// GetAllPosts fetches all blog posts
	GetAllPosts(ctx context.Context) (*blog.Blog, error)
}

// blogService implements the BlogService interface
type blogService struct {
	postRepository blog.PostRepository
}

// NewBlogService creates a new BlogService instance
func NewBlogService(postRepository blog.PostRepository) BlogService {
	return &blogService{
		postRepository: postRepository,
	}
}

// GetAllPosts fetches all blog posts and sorts them by filename
func (s *blogService) GetAllPosts(ctx context.Context) (*blog.Blog, error) {
	posts, err := s.postRepository.FetchPosts(ctx)
	if err != nil {
		return nil, err
	}

	// Sort posts by filename in descending order (newest first)
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Filename > posts[j].Filename
	})

	return &blog.Blog{
		Posts: posts,
	}, nil
}
