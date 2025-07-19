package main

import (
	"net/http"
	"strconv"

	"github.com/MatiasRoje/go-with-vue/backend/internal/models"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			returnErrorResponse(c, err.Error())
			return
		}

		// TODO: Add login logic

		c.IndentedJSON(http.StatusOK, LoginResponse{Error: false, Message: "Login successful"})
	}
}

// User handlers
type getUserResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

type getUsersResponse struct {
	Error   bool          `json:"error"`
	Message string        `json:"message"`
	Users   []models.User `json:"users"`
}

func getUsersHandler(app *application) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Query("email")
		if email != "" {
			user, err := app.models.Users.GetByEmail(email)
			if err != nil {
				returnErrorResponse(c, err.Error())
				return
			}
			c.IndentedJSON(http.StatusOK, getUserResponse{Error: false, Message: "User retrieved successfully", User: user})
			return
		}

		users, err := app.models.Users.GetAll()
		if err != nil {
			returnErrorResponse(c, err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, getUsersResponse{Error: false, Message: "Users retrieved successfully", Users: users})
	}
}

func getUserHandler(app *application) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			returnErrorResponse(c, err.Error())
			return
		}
		user, err := app.models.Users.GetByID(id)
		if err != nil {
			returnErrorResponse(c, err.Error())
			return
		}

		// NOTE: Maybe we shouldn't return the password?
		c.IndentedJSON(http.StatusOK, getUserResponse{Error: false, Message: "User retrieved successfully", User: user})
	}
}

type createUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type createUserResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func createUserHandler(app *application) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			returnErrorResponse(c, err.Error())
			return
		}

		err := app.models.Users.Insert(models.User{
			Email:     req.Email,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Password:  req.Password,
		})
		if err != nil {
			returnErrorResponse(c, err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, createUserResponse{Error: false, Message: "User created successfully"})
	}
}
