package fetcher

type BlogFetcher struct {
	BlogProvider Fetcher
}

func (bf *BlogFetcher) Fetch() *Blog {
	return bf.BlogProvider.Fetch()
}

type Fetcher interface {
	Fetch() *Blog
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
