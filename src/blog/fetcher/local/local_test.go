package local

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalBlogFetcher_Fetch(t *testing.T) {
	l := &LocalBlogFetcher{
		Context: &gin.Context{},
	}

	blog := l.Fetch()
	assert.True(t, len(blog.Posts) > 0)
}
