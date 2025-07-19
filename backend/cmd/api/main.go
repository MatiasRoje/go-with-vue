package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	// Server
	port int

	// database
	db_host     string
	db_port     string
	db_user     string
	db_password string
	db_name     string
}

type application struct {
	config config
	router *gin.Engine
	db     *sql.DB
}

func main() {
	cfg := config{
		port: 3001,

		// database
		db_host:     os.Getenv("DB_HOST"),
		db_port:     os.Getenv("DB_PORT"),
		db_user:     os.Getenv("DB_USER"),
		db_password: os.Getenv("DB_PASSWORD"),
		db_name:     os.Getenv("DB_NAME"),
	}

	db, err := initDB(cfg)
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		router: gin.Default(),
		db:     db,
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

	api.POST("/login", LoginHandler)

	app.router.Run(fmt.Sprintf(":%d", app.config.port))
}
