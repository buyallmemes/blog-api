package local

import (
	md "buyallmemes.com/blog/md/bmparser"
	"buyallmemes/blog/fetcher"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type LocalBlogFetcher struct {
}

const postsLocation = "/posts/"

func (local *LocalBlogFetcher) Fetch() *fetcher.Blog {
	blog := fetcher.NewBlog()
	postsAbsolutePath := constructAbsolutePath()
	dir, _ := os.ReadDir(postsAbsolutePath)
	channels := make([]chan fetcher.Post, 0)
	for _, file := range dir {
		if !strings.HasSuffix(file.Name(), ".md") {
			continue
		}
		outputChannel := local.fetchPost(postsAbsolutePath, file)
		channels = append(channels, outputChannel)
	}

	for _, ch := range channels {
		blog.Posts = append(blog.Posts, <-ch)
	}

	return blog
}

func constructAbsolutePath() string {
	currentDir, _ := os.Getwd()
	projectDir := strings.Index(currentDir, "/src/")
	postsAbsolutePath := currentDir[:projectDir] + postsLocation
	log.Println(postsAbsolutePath)
	return postsAbsolutePath
}

func (local *LocalBlogFetcher) fetchPost(path string, file os.DirEntry) chan fetcher.Post {
	c := make(chan fetcher.Post)
	go func() {
		content := local.GetPostContent(path, file)
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

func (local *LocalBlogFetcher) GetPostContent(path string, file os.DirEntry) string {
	filePath := filepath.Join(path, file.Name())
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(fileContent)

}

func New() *LocalBlogFetcher {
	return &LocalBlogFetcher{}
}
