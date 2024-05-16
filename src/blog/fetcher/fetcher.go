package fetcher

import (
	md "buyallmemes.com/blog/md/bmparser"
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

func GetPosts(context *gin.Context) {
	blog := fetchBlog(context)
	context.IndentedJSON(http.StatusOK, blog)
}

func fetchBlog(context *gin.Context) *Blog {
	client := github.NewClient(nil)
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		client.WithAuthToken(token)
	}

	blog := Blog{
		Posts: []Post{},
	}
	_, directoryContent, resp, err := client.Repositories.GetContents(context, "buyallmemes", "blog-api", "posts", nil)

	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal(fmt.Printf("Non 200 reponse code while requesting repository: %v", resp))
	}

	slices.Reverse(directoryContent)
	for _, file := range directoryContent {
		if !strings.HasSuffix(*file.Name, ".md") {
			continue
		}
		content, _, response, err := client.Repositories.GetContents(context, "buyallmemes", "blog-api", "posts/"+*file.Name, nil)
		if err != nil {
			log.Fatal(err)
		}
		if response.StatusCode != http.StatusOK {
			log.Fatal(fmt.Printf("Non 200 reponse code while requesting file: %v", response))
		}
		decoded, err := base64.StdEncoding.DecodeString(*content.Content)
		if err != nil {
			log.Fatal(err)
		}

		parserMD := md.ToHTML(string(decoded))

		post := Post{
			Filename: *file.Name,
			Content:  parserMD.Html,
			Date:     parserMD.Date,
			Title:    parserMD.Title,
			Anchor:   parserMD.Anchor,
		}
		blog.Posts = append(blog.Posts, post)
	}

	return &blog
}

type Post struct {
	Filename string `json:"filename"`
	Content  string `json:"content"`
	Date     string `json:"date"`
	Title    string `json:"title"`
	Anchor   string `json:"anchor"`
}

type Blog struct {
	Posts []Post `json:"posts"`
}
