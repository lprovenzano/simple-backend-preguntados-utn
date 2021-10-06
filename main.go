package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type question struct {
	ID       int64    `json:"id"`
	Title    string   `json:"title"`
	Options  []answer `json:"options"`
	UrlImage string   `json:"url_image"`
}

type answer struct {
	Title   string `json:"title"`
	IsRight bool   `json:"right"`
}

var questions = []question{
	{ID: 1,
		Title: "Blue Train", Options: []answer{
		{
			Title:   "Test 1",
			IsRight: true,
		}, {
			Title:   "Test 2 ",
			IsRight: false,
		}, {
			Title:   "Test 3",
			IsRight: false,
		}}, UrlImage: "http://dummyimage.com",
	}, {
		ID: 2,
		Title: "Blue Car", Options: []answer{
		{
			Title:   "Test 1",
			IsRight: false,
		}, {
			Title:   "Test 2 ",
			IsRight: true,
		}, {
			Title:   "Test 3",
			IsRight: false,
		}}, UrlImage: "http://dummyimage.com",
	}, {
		ID: 3,
		Title: "Blue Truck", Options: []answer{
		{
			Title:   "Test 1",
			IsRight: false,
		}, {
			Title:   "Test 2 ",
			IsRight: false,
		}, {
			Title:   "Test 3",
			IsRight: true,
		}}, UrlImage: "http://dummyimage.com",
	},
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	router.GET("/", getQuestions)

	router.Run(fmt.Sprintf(":%s", port))
}

func getQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, questions)
}
