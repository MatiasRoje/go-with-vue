package main

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

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

func returnErrorResponse(c *gin.Context, message string, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.IndentedJSON(statusCode, StandardAPIResponse{Error: true, Message: message})
}

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

	user, err := h.app.models.DBUsers.GetByEmail(req.Email)
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

	frontendUser := userToFrontendUser(*user)

	c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "Login successful", Data: envelope{"token": tokenString, "user": frontendUser}})
}

func (h *Handler) validateTokenHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		returnErrorResponse(c, "No token provided", http.StatusUnauthorized)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := validateJWTToken(tokenString, h.app.config.JwtSecret)
	if err != nil {
		returnErrorResponse(c, "Invalid token", http.StatusUnauthorized)
		return
	}

	user, err := h.app.models.DBUsers.GetByEmail(claims.Email)
	if err != nil {
		returnErrorResponse(c, "Invalid token", http.StatusUnauthorized)
		return
	}

	frontendUser := userToFrontendUser(*user)

	c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "Token is valid", Data: envelope{"user": frontendUser}})
}

// User handlers
func (h *Handler) getUsersHandler(c *gin.Context) {
	email := c.Query("email")
	if email != "" {
		user, err := h.app.models.DBUsers.GetByEmail(email)
		if err != nil {
			returnErrorResponse(c, err.Error())
			return
		}
		c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "User retrieved successfully", Data: envelope{"user": user}})
		return
	}

	users, err := h.app.models.DBUsers.GetAll()
	if err != nil {
		returnErrorResponse(c, err.Error())
		return
	}

	var frontendUsers []frontendUser
	for _, user := range users {
		frontendUsers = append(frontendUsers, userToFrontendUser(*user))
	}

	c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "Users retrieved successfully", Data: envelope{"users": frontendUsers}})
}

func (h *Handler) getUserHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		returnErrorResponse(c, err.Error())
		return
	}
	user, err := h.app.models.DBUsers.GetByID(id)
	if err != nil {
		returnErrorResponse(c, err.Error())
		return
	}

	frontendUser := userToFrontendUser(*user)

	c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "User retrieved successfully", Data: envelope{"user": frontendUser}})
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

	err := h.app.models.DBUsers.Insert(&models.User{
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

// Books handlers
// NOTE: For now we are not using a frontendBook struct, as books don't hold the same type of private fields as users, but we might do it later
func (h *Handler) getBooksHandler(c *gin.Context) {
	books, err := h.app.models.DBBooks.GetAll()
	if err != nil {
		returnErrorResponse(c, err.Error())
		return
	}

	sort.Slice(books, func(i, j int) bool {
		return books[i].ID < books[j].ID
	})

	c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "Books retrieved successfully", Data: envelope{"books": books}})
}

func (h *Handler) getBookHandler(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		returnErrorResponse(c, "No slug provided")
		return
	}
	book, err := h.app.models.DBBooks.GetBySlug(slug)
	if err != nil {
		returnErrorResponse(c, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, StandardAPIResponse{Error: false, Message: "Book retrieved successfully", Data: envelope{"book": book}})
}
