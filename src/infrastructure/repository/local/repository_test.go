package local

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"buyallmemes.com/blog-api/src/domain/blog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for the markdown parser
type MockMarkdownParser struct {
	mock.Mock
}

func (m *MockMarkdownParser) ParseMarkdown(content string) (blog.ParsedMarkdown, error) {
	args := m.Called(content)
	return args.Get(0).(blog.ParsedMarkdown), args.Error(1)
}

func TestLocalRepository_FetchPosts(t *testing.T) {
	// Create a temporary directory for test files
	tempDir := t.TempDir()

	// Create test markdown files
	testFile1 := filepath.Join(tempDir, "test1.md")
	testFile2 := filepath.Join(tempDir, "test2.md")

	content1 := "---\ntitle: Test Post 1\ndate: 2025-05-01\n---\nThis is test content 1"
	content2 := "---\ntitle: Test Post 2\ndate: 2025-05-02\n---\nThis is test content 2"

	err := os.WriteFile(testFile1, []byte(content1), 0644)
	assert.NoError(t, err)
	err = os.WriteFile(testFile2, []byte(content2), 0644)
	assert.NoError(t, err)

	// Create a non-markdown file that should be ignored
	nonMdFile := filepath.Join(tempDir, "notapost.txt")
	err = os.WriteFile(nonMdFile, []byte("not a markdown file"), 0644)
	assert.NoError(t, err)

	// Setup mock parser
	mockParser := new(MockMarkdownParser)
	mockParser.On("ParseMarkdown", content1).Return(blog.ParsedMarkdown{
		Title:   "Test Post 1",
		Content: "This is test content 1",
		Date:    "2025-05-01",
		Anchor:  "test-post-1",
	}, nil)
	mockParser.On("ParseMarkdown", content2).Return(blog.ParsedMarkdown{
		Title:   "Test Post 2",
		Content: "This is test content 2",
		Date:    "2025-05-02",
		Anchor:  "test-post-2",
	}, nil)

	// Create repository with the temp directory
	repo := NewLocalRepository(mockParser, tempDir)

	// Test FetchPosts
	posts, err := repo.FetchPosts(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(posts))

	// Verify that posts contain the expected data
	// Note: order might be different based on filesystem
	postTitles := make(map[string]bool)
	for _, post := range posts {
		postTitles[post.Title] = true
	}
	assert.True(t, postTitles["Test Post 1"])
	assert.True(t, postTitles["Test Post 2"])

	mockParser.AssertExpectations(t)
}

func TestLocalRepository_FetchPosts_DirectoryError(t *testing.T) {
	mockParser := new(MockMarkdownParser)

	// Create repository with a non-existent directory
	repo := NewLocalRepository(mockParser, "/nonexistent/directory")

	// Test FetchPosts should return an error
	posts, err := repo.FetchPosts(context.Background())
	assert.Error(t, err)
	assert.Nil(t, posts)
	assert.Contains(t, err.Error(), "error reading posts directory")
}
