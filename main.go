package main

import (
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/markdown", func(c *gin.Context) {
		body, ok := c.GetPostForm("body")
		if !ok {
			c.JSON(http.StatusBadRequest, "badrequest")
			return
		}
		markdown := blackfriday.MarkdownCommon([]byte(body))
		c.Data(http.StatusOK, "text/html; charset=utf-8", markdown)
	})

	router.Run(":" + port)
}
