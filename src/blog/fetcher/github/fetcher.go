package github

import (
	"buyallmemes.com/blog/md/bmparser"
	"buyallmemes/blog/fetcher"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v62/github"
	"log"
	"net/http"
	"os"
	"slices"
	"strings"
)

func New(context *gin.Context) *BlogFetcher {
	return &BlogFetcher{
		Context: context,
		Client:  newGitHubClientWithToken(),
	}
}

type BlogFetcher struct {
	Context *gin.Context
	Client  *github.Client
}

func (gh *BlogFetcher) Fetch() *fetcher.Blog {
	blog := fetcher.NewBlog()
	directoryContent := gh.GetDirectoryContent()

	slices.Reverse(directoryContent)

	for _, file := range directoryContent {
		if !strings.HasSuffix(*file.Name, ".md") {
			continue
		}
		post := gh.fetchPost(file)
		blog.Posts = append(blog.Posts, post)
	}

	return blog
}

func (gh *BlogFetcher) fetchPost(file *github.RepositoryContent) fetcher.Post {
	content := gh.GetPostContent(file.Name)
	decoded, err := base64.StdEncoding.DecodeString(*content.Content)
	if err != nil {
		log.Fatal(err)
	}

	parserMD := md.ToHTML(string(decoded))

	post := fetcher.Post{
		Filename: *file.Name,
		Content:  parserMD.Html,
		Date:     parserMD.Date,
		Title:    parserMD.Title,
		Anchor:   parserMD.Anchor,
	}
	return post
}

func newGitHubClientWithToken() *github.Client {
	client := github.NewClient(nil)
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		client.WithAuthToken(token)
	}
	return client
}

func (gh *BlogFetcher) GetDirectoryContent() []*github.RepositoryContent {
	_, directoryContent, resp, err := gh.Client.Repositories.GetContents(gh.Context, "buyallmemes", "blog-api", "posts", nil)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(fmt.Printf("Non 200 reponse code while requesting repository: %v", resp))
	}
	return directoryContent
}

func (gh *BlogFetcher) GetPostContent(filename *string) *github.RepositoryContent {
	content, _, response, err := gh.Client.Repositories.GetContents(gh.Context, "buyallmemes", "blog-api", "posts/"+*filename, nil)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		log.Fatal(fmt.Printf("Non 200 reponse code while requesting file: %v", response))
	}
	return content
}
