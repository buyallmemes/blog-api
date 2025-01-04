package main

import (
	"buyallmemes.com/blog-api/src/blog/fetcher"
	"buyallmemes.com/blog-api/src/blog/fetcher/local"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := setupEngine()
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func setupEngine() *gin.Engine {
	engine := gin.Default()
	engine.GET("/posts", getPosts)
	return engine
}

func getPosts(context *gin.Context) {
	backend := local.New()

	blogFetcher := fetcher.BlogFetcher{
		BlogProvider: backend,
	}
	blog := blogFetcher.Fetch()
	context.JSON(http.StatusOK, blog)
	errors := context.Errors
	for _, err := range errors {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
}
