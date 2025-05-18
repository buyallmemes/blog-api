package blog

// MarkdownParser defines the interface for parsing markdown content
type MarkdownParser interface {
	// ParseMarkdown parses markdown content and returns parsed markdown data
	ParseMarkdown(content string) (ParsedMarkdown, error)
}
