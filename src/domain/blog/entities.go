package blog

// Post represents a blog post
type Post struct {
	Filename string
	Content  string
	Date     string
	Title    string
	Anchor   string
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
	Posts []Post
}

// NewBlog creates a new Blog instance
func NewBlog() *Blog {
	return &Blog{
		Posts: []Post{},
	}
}
