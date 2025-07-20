package main

import (
	"net/http"
	"strconv"

	"github.com/MatiasRoje/go-with-vue/backend/internal/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	app *application
}

type StandardAPIResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type envelope map[string]any

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		returnErrorResponse(c, err.Error())
		return
	}

	user, err := h.app.models.Users.GetByEmail(req.Email)
	if err != nil {
		returnErrorResponse(c, "Invalid credentials")
		return
	}

	if err := user.CheckPassword(req.Password); err != nil {
		returnErrorResponse(c, "Invalid credentials")
		return
	}

	tokenString, err := generateJWTToken(user.Email, h.app.config.JwtSecret)
	if err != nil {
		returnErrorResponse(c, "error generating JWT token")
		return
	}

	c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "Password is correct", Data: envelope{"token": tokenString}})
}

// User handlers
func (h *Handler) getUsersHandler(c *gin.Context) {
	email := c.Query("email")
	if email != "" {
		user, err := h.app.models.Users.GetByEmail(email)
		if err != nil {
			returnErrorResponse(c, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "User retrieved successfully", Data: envelope{"user": user}})
		return
	}

	users, err := h.app.models.Users.GetAll()
	if err != nil {
		returnErrorResponse(c, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "Users retrieved successfully", Data: envelope{"users": users}})
}

func (h *Handler) getUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		returnErrorResponse(c, err.Error())
		return
	}
	user, err := h.app.models.Users.GetByID(id)
	if err != nil {
		returnErrorResponse(c, err.Error())
		return
	}

	// TODO: Maybe we shouldn't return the password?
	c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "User retrieved successfully", Data: envelope{"user": user}})
}

type createUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func (h *Handler) createUserHandler(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		returnErrorResponse(c, err.Error())
		return
	}

	err := h.app.models.Users.Insert(models.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	})
	if err != nil {
		returnErrorResponse(c, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "User created successfully"})
}
