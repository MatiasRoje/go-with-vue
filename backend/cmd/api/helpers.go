package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func returnErrorResponse(c *gin.Context, message string, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.IndentedJSON(statusCode, LoginResponse{Error: true, Message: message})
}
