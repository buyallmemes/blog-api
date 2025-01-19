package fetcher

import (
	"context"
	"sort"
)

type BlogFetcher struct {
	BlogProvider Fetcher
}

func (bf *BlogFetcher) Fetch(ctx context.Context) *Blog {
	blog := bf.BlogProvider.Fetch(ctx)
	sortByFilename(blog.Posts)
	return blog
}

func sortByFilename(posts []Post) {
	sort.Slice(posts, func(post1, post2 int) bool {
		return posts[post1].Filename > posts[post2].Filename
	})
}

type Fetcher interface {
	Fetch(ctx context.Context) *Blog
}

type Blog struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
	Date     string `json:"date"`
	Title    string `json:"title"`
	Anchor   string `json:"anchor"`
}

func NewBlog() *Blog {
	return &Blog{
		Posts: []Post{},
	}
}
