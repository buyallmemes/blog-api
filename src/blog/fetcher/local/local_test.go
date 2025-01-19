package local

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalBlogFetcher_Fetch(t *testing.T) {
	l := New()

	blog := l.Fetch(context.Background())
	assert.True(t, len(blog.Posts) > 0)
}
