package router

import (
	"jabutech/blog-restful-api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(pingHandler handler.PingHandler) *gin.Engine {
	// Create router with gin
	router := gin.Default()
	// Router group api
	api := router.Group("/api")

	// Endpoint ping
	api.GET("/ping", pingHandler.Ping)

	return router
}
