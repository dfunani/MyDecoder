package routes

import (
	"dfunani/mydecoder/routes/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PageRoutes(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/pages/encode")
	})
	router.GET("/encode", handlers.HandleHomePage)
	router.GET("/decode", handlers.HandleHomePage)
}
