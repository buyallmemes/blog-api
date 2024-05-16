package main

import (
	"buyallmemes/blog/fetcher"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	engine.GET("/posts", fetcher.GetPosts)
	err := engine.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
