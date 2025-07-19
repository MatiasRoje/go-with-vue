package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, LoginResponse{Error: true, Message: err.Error()})
		return
	}

	// TODO: Add login logic
	
	c.IndentedJSON(http.StatusOK, LoginResponse{Error: false, Message: "Login successful"})
}