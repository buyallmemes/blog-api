package blog

import (
	"encoding/json"
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

func TestBlogSerialization(t *testing.T) {
	// Test with a blog containing posts
	blog := Blog{
		Posts: []Post{
			{
				Filename: "test.md",
				Content:  "<p>Test content</p>",
				Date:     "2023-01-01",
				Title:    "Test Post",
				Anchor:   "test-post",
			},
		},
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(blog)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Deserialize from JSON
	var deserializedBlog Blog
	err = json.Unmarshal(jsonData, &deserializedBlog)
	assert.NoError(t, err)
	assert.Len(t, deserializedBlog.Posts, 1)
	assert.Equal(t, "Test Post", deserializedBlog.Posts[0].Title)
}

func TestEmptyBlogSerialization(t *testing.T) {
	// Test with an empty blog
	blog := Blog{
		Posts: []Post{},
	}

	// Serialize to JSON
	jsonData, err := json.Marshal(blog)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Check the JSON structure
	jsonString := string(jsonData)
	assert.Contains(t, jsonString, "\"posts\":[]")

	// Deserialize from JSON
	var deserializedBlog Blog
	err = json.Unmarshal(jsonData, &deserializedBlog)
	assert.NoError(t, err)
	assert.NotNil(t, deserializedBlog.Posts)
	assert.Empty(t, deserializedBlog.Posts)
}

func TestNilBlogSerialization(t *testing.T) {
	// Test with a blog with nil Posts
	blog := Blog{}

	// Serialize to JSON
	jsonData, err := json.Marshal(blog)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Check the JSON structure - should be empty array, not null
	jsonString := string(jsonData)
	assert.Contains(t, jsonString, "\"posts\":[]")
	assert.NotContains(t, jsonString, "\"posts\":null")

	// Deserialize from JSON
	var deserializedBlog Blog
	err = json.Unmarshal(jsonData, &deserializedBlog)
	assert.NoError(t, err)

	// Posts should be initialized as an empty slice, not nil
	assert.NotNil(t, deserializedBlog.Posts)
	assert.Empty(t, deserializedBlog.Posts)
}
