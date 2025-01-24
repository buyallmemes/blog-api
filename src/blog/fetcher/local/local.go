package local

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"buyallmemes.com/blog-api/src/blog/fetcher"
	"buyallmemes.com/blog-api/src/blog/md"
)

type LocalBlogFetcher struct {
}

const postsLocation = "/posts/"

func New() *LocalBlogFetcher {
	return &LocalBlogFetcher{}
}

func (local *LocalBlogFetcher) Fetch(_ context.Context) *fetcher.Blog {
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
	rootPath, err := getModuleRootPath()
	if err != nil {
		log.Fatal(err)
	}
	postsAbsolutePath := rootPath + postsLocation
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

func getModuleRootPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("go.mod not found")
}
