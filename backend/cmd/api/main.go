package main

import (
	"fmt"
	"log"

	"github.com/MatiasRoje/go-with-vue/backend/internal/config"
	"github.com/MatiasRoje/go-with-vue/backend/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type application struct {
	config config.Config
	router *gin.Engine
	models *database.Models
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	db, err := database.InitDB(*cfg)
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
	defer db.Close()

	app := &application{
		config: *cfg,
		router: gin.Default(),
		models: database.NewDBModels(db),
	}

	app.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Content-Type", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Initialize handler and routes
	h := &Handler{app: app}
	api := app.router.Group("/api/v1")

	// Register routes
	api.POST("/login", h.LoginHandler)
	api.POST("/validate-token", h.validateTokenHandler)
	api.GET("/users", h.getUsersHandler)
	api.GET("/users/:id", h.getUserHandler)
	api.POST("/users", h.createUserHandler)

	app.router.Run(fmt.Sprintf(":%s", app.config.AppPort))
}
