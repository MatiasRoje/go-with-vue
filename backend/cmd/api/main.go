package main

import (
	"fmt"
	"log"
	"net/http"

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

	// Api group
	api := app.router.Group("/api/v1")

	// Register routes
	app.router.GET("/", func(c *gin.Context) {
		var payload struct {
			Okay    bool   `json:"okay"`
			Message string `json:"message"`
		}
		payload.Okay = true
		payload.Message = "Hello, world"

		c.IndentedJSON(http.StatusOK, payload)
	})

	api.POST("/login", LoginHandler())
	api.GET("/users", getUsersHandler(app))
	api.GET("/users/:id", getUserHandler(app))
	api.POST("/users", createUserHandler(app))

	app.router.Run(fmt.Sprintf(":%s", app.config.AppPort))
}
