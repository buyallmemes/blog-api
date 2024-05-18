package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getPosts(t *testing.T) {
	engine := setupEngine()

	w := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/posts", nil)
	engine.ServeHTTP(w, request)
	log.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)

}
