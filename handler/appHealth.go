package handler

import "github.com/gin-gonic/gin"

type appHealthHandler struct {
}

// Instance
func NewAppHealthHandler() *appHealthHandler {
	return &appHealthHandler{}
}

// Function check app health
func (h *appHealthHandler) Check(c *gin.Context) {
	c.JSON(200, "I am healthy.")
}
