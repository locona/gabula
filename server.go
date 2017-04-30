package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func main() {
	router := router.Default()

	router.LoadHTMLGlob("src/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"src/templates/index.html",
			gin.H{
				"title": "Home Page",
			},
		)

	})

}
