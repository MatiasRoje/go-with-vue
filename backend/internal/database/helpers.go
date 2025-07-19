package database

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("go123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error generating password hash: %w", err)
	}

	// Insert the test user
	query := `
		INSERT INTO users (email, first_name, last_name, password)
		VALUES ('go@vue.com', 'Go', 'Vue', $1)
	`

	if _, err := db.Exec(query, hashedPassword); err != nil {
		return fmt.Errorf("error inserting test user: %w", err)
	}

	log.Println("Test user successfully inserted")
	return nil
}

func createTokensTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS tokens (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			user_email VARCHAR(255) UNIQUE NOT NULL,
			token VARCHAR(255) NOT NULL,
			token_hash BYTEA NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			expires_at TIMESTAMP NOT NULL
		)
	`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("error creating tokens table: %w", err)
	}

	log.Println("Tokens table successfully initialized")
	return nil
}
