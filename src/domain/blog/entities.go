package blog

import (
	"encoding/json"
)

// Post represents a blog post
type Post struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
	Date     string `json:"date"`
	Title    string `json:"title"`
	Anchor   string `json:"anchor"`
}

// ParsedMarkdown represents the result of parsing markdown content
type ParsedMarkdown struct {
	Content string
	Title   string
	Date    string
	Anchor  string
}

// Blog represents a collection of blog posts
type Blog struct {
	Posts []Post `json:"posts"`
}

// NewBlog creates a new Blog instance
func NewBlog() *Blog {
	return &Blog{
		Posts: []Post{},
	}
}

// MarshalJSON implements the json.Marshaler interface to ensure Posts is never null in JSON
func (b Blog) MarshalJSON() ([]byte, error) {
	type Alias Blog
	return json.Marshal(&struct {
		Posts []Post `json:"posts"`
		*Alias
	}{
		Posts: func() []Post {
			if b.Posts == nil {
				return []Post{}
			}
			return b.Posts
		}(),
		Alias: (*Alias)(&b),
	})
}
