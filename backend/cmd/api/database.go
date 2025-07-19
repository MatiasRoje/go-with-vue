package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	DB_MAX_OPEN_CONNS = 5
	DB_MAX_IDLE_CONNS = 5
	DB_MAX_LIFETIME   = 5 * time.Minute
)

func connectToDB(cfg config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5", cfg.db_host, cfg.db_port, cfg.db_user, cfg.db_password, cfg.db_name)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	db.SetMaxOpenConns(DB_MAX_OPEN_CONNS)
	db.SetMaxIdleConns(DB_MAX_IDLE_CONNS)
	db.SetConnMaxLifetime(DB_MAX_LIFETIME)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db: %w", err)
	}

	return db, nil
}

func createUsersTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			password VARCHAR(80) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("error creating users table: %w", err)
	}

	log.Println("Users table successfully initialized")
	return nil
}

func insertTestUser(db *sql.DB) error {
	// Check whether the test user already exists
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = 'go@vue.com'").Scan(&count); err != nil {
		return fmt.Errorf("error checking for test user: %w", err)
	}

	if count > 0 {
		log.Println("Test user already exists")
		return nil
	}

	// Insert the test user
	query := `
		INSERT INTO users (email, first_name, last_name, password)
		VALUES ('go@vue.com', 'Go', 'Vue', 'go123')
	`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("error inserting test user: %w", err)
	}

	log.Println("Test user successfully inserted")
	return nil
}

func initDB(cfg config) (*sql.DB, error) {
	db, err := connectToDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	if err := createUsersTable(db); err != nil {
		return nil, fmt.Errorf("error creating users table: %w", err)
	}

	if err := insertTestUser(db); err != nil {
		return nil, fmt.Errorf("error inserting test user: %w", err)
	}

	return db, nil
}
