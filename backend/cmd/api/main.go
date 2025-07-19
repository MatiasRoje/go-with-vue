package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type config struct {
	port int
}

type application struct {
	config config	
	router *gin.Engine
}	

func main() {
	cfg := config{
		port: 3001,
	}

	app := &application{
		config: cfg,
		router: gin.Default(),
	}
	app.router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:3000"},
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Accept", "Content-Type", "Authorization", "X-CSRF-Token"},
    ExposeHeaders:    []string{"Link"},
    AllowCredentials: true,
    MaxAge: 300,
  }))

	// Api group
	api := app.router.Group("/api/v1")
	
	// Register routes
	app.router.GET("/", func(c *gin.Context) {
		var payload struct {
			Okay bool `json:"okay"`
			Message string `json:"message"`
		}
		payload.Okay = true
		payload.Message = "Hello, world"

		c.JSON(http.StatusOK, payload)
	})

	api.POST("/login", LoginHandler)

	app.router.Run(fmt.Sprintf(":%d", app.config.port))
}