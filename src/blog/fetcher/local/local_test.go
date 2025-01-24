package local

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalBlogFetcher_Fetch(t *testing.T) {
	l := New()

	blog := l.Fetch(context.Background())
	assert.True(t, len(blog.Posts) > 0)
}
