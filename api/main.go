package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("pages/*")
	router.Static("/assets", "./assets/")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "EstefDev - Programador Junior",
		})
	})
	router.Run()
}
