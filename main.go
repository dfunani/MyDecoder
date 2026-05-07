package main

import (
	"dfunani/mydecoder/routes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Decoder Started")

	client := gin.Default()
	client.LoadHTMLGlob("templates/*")

	client.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/pages/encode")
	})

	apis := client.Group("/api")
	static := client.Group("/pages")
	routes.ApiRoutes(apis)
	routes.PageRoutes(static)

	client.Run(":80")
}
