package routes

import (
	"dfunani/mydecoder/routes/handlers"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.RouterGroup) {
	router.POST("/encode", handlers.HandleEncodeMessage)
	router.POST("/decode", handlers.HandleDecodeMessage)
}
