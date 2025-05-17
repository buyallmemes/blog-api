package markdown

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"buyallmemes.com/blog-api/src/domain/blog"
	"github.com/gosimple/slug"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

// Common errors
var (
	ErrEmptyMarkdown       = errors.New("empty markdown content")
	ErrMarkdownConversion  = errors.New("markdown conversion failed")
	ErrFrontmatterDecoding = errors.New("frontmatter decoding failed")
)

// Frontmatter metadata structure
type frontmatterMeta struct {
	Title string `yaml:"title"`
	Date  string `yaml:"date"`
}

// GoldmarkParser implements the MarkdownParser interface using the Goldmark library
type GoldmarkParser struct {
	markdown goldmark.Markdown
	slugify  func(string) string
}

// NewGoldmarkParser creates a new GoldmarkParser instance
func NewGoldmarkParser() *GoldmarkParser {
	return &GoldmarkParser{
		markdown: goldmark.New(
			goldmark.WithExtensions(
				&frontmatter.Extender{},
			),
		),
		slugify: slug.Make,
	}
}

// ParseMarkdown parses markdown content and returns parsed markdown data
func (p *GoldmarkParser) ParseMarkdown(source string) (blog.ParsedMarkdown, error) {
	// Validate input
	if strings.TrimSpace(source) == "" {
		return blog.ParsedMarkdown{}, ErrEmptyMarkdown
	}

	var buf bytes.Buffer
	context := parser.NewContext()
	withContext := parser.WithContext(context)

	// Convert markdown to HTML
	err := p.markdown.Convert([]byte(source), &buf, withContext)
	if err != nil {
		return blog.ParsedMarkdown{}, fmt.Errorf("%w: %v", ErrMarkdownConversion, err)
	}

	// Initialize result with HTML content
	result := blog.ParsedMarkdown{
		Content: buf.String(),
	}

	// Extract and process frontmatter
	if d := frontmatter.Get(context); d != nil {
		meta := frontmatterMeta{}

		if err := d.Decode(&meta); err != nil {
			// Return partial result with content but no metadata
			return blog.ParsedMarkdown{Content: result.Content}, fmt.Errorf("%w: %v", ErrFrontmatterDecoding, err)
		}

		// Set metadata fields
		result.Title = meta.Title
		result.Date = meta.Date

		// Generate anchor from title if available
		if result.Title != "" {
			result.Anchor = p.slugify(result.Title)
		}
	}

	return result, nil
}
