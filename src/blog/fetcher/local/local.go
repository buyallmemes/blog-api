package local

import (
	md "buyallmemes.com/blog/md/bmparser"
	"buyallmemes/blog/fetcher"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type LocalBlogFetcher struct {
	Context *gin.Context
}

const postsLocation = "/Users/mark/dev/private/buyallmemes/blog-api/posts"

func (local *LocalBlogFetcher) Fetch() *fetcher.Blog {
	blog := fetcher.NewBlog()

	dir, _ := os.ReadDir(postsLocation)
	channels := make([]chan fetcher.Post, 0)
	for _, file := range dir {
		if !strings.HasSuffix(file.Name(), ".md") {
			continue
		}
		outputChannel := local.fetchPost(file)
		channels = append(channels, outputChannel)
	}

	for _, ch := range channels {
		blog.Posts = append(blog.Posts, <-ch)
	}

	return blog
}

func (local *LocalBlogFetcher) fetchPost(file os.DirEntry) chan fetcher.Post {
	c := make(chan fetcher.Post)
	go func() {
		content := local.GetPostContent(file)
		parserMD := md.ToHTML(content)

		filName := file.Name()
		post := fetcher.Post{
			Filename: filName,
			Content:  parserMD.Html,
			Date:     parserMD.Date,
			Title:    parserMD.Title,
			Anchor:   parserMD.Anchor,
		}
		c <- post
	}()
	return c
}

func (local *LocalBlogFetcher) GetPostContent(file os.DirEntry) string {
	filePath := filepath.Join(postsLocation, file.Name())
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(fileContent)

}

func New(context *gin.Context) *LocalBlogFetcher {
	return &LocalBlogFetcher{
		Context: context,
	}
}
