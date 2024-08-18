package main

import (
	"buyallmemes.com/blog-api/pkg/blog/fetcher"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getPosts(t *testing.T) {
	engine := setupEngine()

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/posts", nil)
	engine.ServeHTTP(w, request)
	assert.Equal(t, 200, w.Code)
	body, err := io.ReadAll(w.Body)
	assert.NoError(t, err)

	blog := new(fetcher.Blog)
	unmarshallingError := json.Unmarshal(body, &blog)
	assert.NoError(t, unmarshallingError)
	posts := blog.Posts
	assert.True(t, len(posts) > 0)
	assert.Equal(t, "Hello, World!", posts[len(posts)-1].Title)
}
