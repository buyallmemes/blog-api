package blog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlog(t *testing.T) {
	blog := NewBlog()

	assert.NotNil(t, blog)
	assert.Empty(t, blog.Posts)
}

func TestBlog_AddPost(t *testing.T) {
	blog := NewBlog()
	post := Post{
		Filename: "test.md",
		Content:  "<p>Test content</p>",
		Date:     "2024-05-17",
		Title:    "Test Post",
		Anchor:   "test-post",
	}

	blog.Posts = append(blog.Posts, post)

	assert.Len(t, blog.Posts, 1)
	assert.Equal(t, post, blog.Posts[0])
}
