package local

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalBlogFetcher_Fetch(t *testing.T) {
	l := New()

	blog := l.Fetch()
	assert.True(t, len(blog.Posts) > 0)
}
