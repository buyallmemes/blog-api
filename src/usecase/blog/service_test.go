package blog

import (
	"context"
	"errors"
	"testing"

	"buyallmemes.com/blog-api/src/domain/blog"
	"github.com/stretchr/testify/assert"
)

// StubPostRepository is a stub implementation of the PostRepository interface
type StubPostRepository struct {
	posts []blog.Post
	err   error
}

func (s *StubPostRepository) FetchPosts(ctx context.Context) ([]blog.Post, error) {
	return s.posts, s.err
}

func TestNewBlogService(t *testing.T) {
	repo := &StubPostRepository{}
	service := NewBlogService(repo)

	assert.NotNil(t, service)
}

func TestBlogService_GetAllPosts_Success(t *testing.T) {
	posts := []blog.Post{
		{
			Filename: "test1.md",
			Content:  "<p>Test content 1</p>",
			Date:     "2024-05-17",
			Title:    "Test Post 1",
			Anchor:   "test-post-1",
		},
		{
			Filename: "test2.md",
			Content:  "<p>Test content 2</p>",
			Date:     "2024-05-16",
			Title:    "Test Post 2",
			Anchor:   "test-post-2",
		},
	}

	repo := &StubPostRepository{posts: posts}
	service := NewBlogService(repo)

	result, err := service.GetAllPosts(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Posts, 2)
	// Check that posts are sorted by filename in descending order (newest first)
	assert.Equal(t, "test2.md", result.Posts[0].Filename)
	assert.Equal(t, "test1.md", result.Posts[1].Filename)
}

func TestBlogService_GetAllPosts_Error(t *testing.T) {
	expectedError := errors.New("repository error")
	repo := &StubPostRepository{err: expectedError}
	service := NewBlogService(repo)

	result, err := service.GetAllPosts(context.Background())

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
}
