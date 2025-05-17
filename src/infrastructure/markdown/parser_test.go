package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGoldmarkParser(t *testing.T) {
	parser := NewGoldmarkParser()

	assert.NotNil(t, parser)
	assert.NotNil(t, parser.markdown)
}

func TestGoldmarkParser_ParseMarkdown_Simple(t *testing.T) {
	parser := NewGoldmarkParser()

	parsed, err := parser.ParseMarkdown("Hello, World!")

	assert.NoError(t, err)
	assert.Equal(t, "<p>Hello, World!</p>\n", parsed.Content)
	assert.Empty(t, parsed.Title)
	assert.Empty(t, parsed.Date)
	assert.Empty(t, parsed.Anchor)
}

func TestGoldmarkParser_ParseMarkdown_WithFrontmatter(t *testing.T) {
	parser := NewGoldmarkParser()

	markdown := `---
title: Test Title
date: 2024-05-17
---
Hello, World!`

	parsed, err := parser.ParseMarkdown(markdown)

	assert.NoError(t, err)
	assert.Equal(t, "<p>Hello, World!</p>\n", parsed.Content)
	assert.Equal(t, "Test Title", parsed.Title)
	assert.Equal(t, "2024-05-17", parsed.Date)
	assert.Equal(t, "test-title", parsed.Anchor)
}

func TestGoldmarkParser_ParseMarkdown_WithMarkdownFormatting(t *testing.T) {
	parser := NewGoldmarkParser()

	markdown := `# Heading

This is a paragraph with **bold** and *italic* text.

- List item 1
- List item 2
`

	parsed, err := parser.ParseMarkdown(markdown)

	assert.NoError(t, err)
	assert.Contains(t, parsed.Content, "<h1>Heading</h1>")
	assert.Contains(t, parsed.Content, "<p>This is a paragraph with <strong>bold</strong> and <em>italic</em> text.</p>")
	assert.Contains(t, parsed.Content, "<ul>")
	assert.Contains(t, parsed.Content, "<li>List item 1</li>")
	assert.Contains(t, parsed.Content, "<li>List item 2</li>")
	assert.Empty(t, parsed.Title)
	assert.Empty(t, parsed.Date)
	assert.Empty(t, parsed.Anchor)
}
