package github

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"buyallmemes.com/blog-api/src/blog/fetcher"
	"buyallmemes.com/blog-api/src/blog/md"
	"github.com/google/go-github/v62/github"
	"github.com/mfenderov/konfig"
)

func New() *GithubBlogFetcher {
	return &GithubBlogFetcher{
		Client: newGitHubClientWithToken(),
	}
}

type GithubBlogFetcher struct {
	Client *github.Client
}

func (gh *GithubBlogFetcher) Fetch(ctx context.Context) *fetcher.Blog {
	blog := fetcher.NewBlog()
	directoryContent := gh.GetDirectoryContent(ctx)

	channels := make([]chan fetcher.Post, 0)
	for _, file := range directoryContent {
		if !strings.HasSuffix(*file.Name, ".md") {
			continue
		}
		resultChannel := gh.fetchPost(ctx, file)
		channels = append(channels, resultChannel)
	}

	for _, ch := range channels {
		blog.Posts = append(blog.Posts, <-ch)
	}

	return blog
}

func (gh *GithubBlogFetcher) fetchPost(ctx context.Context, file *github.RepositoryContent) chan fetcher.Post {
	c := make(chan fetcher.Post)
	go func() {
		content := gh.GetPostContent(ctx, file.Name)
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
		c <- post
	}()
	return c
}

func newGitHubClientWithToken() *github.Client {
	client := github.NewClient(nil)
	token := konfig.GetEnv("github.token")
	if token != "" {
		client.WithAuthToken(token)
	}
	return client
}

func (gh *GithubBlogFetcher) GetDirectoryContent(ctx context.Context) []*github.RepositoryContent {
	_, directoryContent, resp, err := gh.Client.Repositories.GetContents(ctx, "buyallmemes", "blog-api", "posts", nil)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(fmt.Printf("Non 200 reponse code while requesting repository: %v", resp))
	}
	return directoryContent
}

func (gh *GithubBlogFetcher) GetPostContent(ctx context.Context, filename *string) *github.RepositoryContent {
	content, _, response, err := gh.Client.Repositories.GetContents(ctx, "buyallmemes", "blog-api", "posts/"+*filename, nil)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		log.Fatal(fmt.Printf("Non 200 reponse code while requesting file: %v", response))
	}
	return content
}
