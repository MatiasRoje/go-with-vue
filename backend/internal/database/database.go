package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MatiasRoje/go-with-vue/backend/internal/config"
	_ "github.com/lib/pq"
)

const (
	DB_MAX_OPEN_CONNS = 5
	DB_MAX_IDLE_CONNS = 5
	DB_MAX_LIFETIME   = 5 * time.Minute
)

func connectToDB(cfg config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
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

func InitDB(cfg config.Config) (*sql.DB, error) {
	db, err := connectToDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	if err := createUsersTable(db); err != nil {
		return nil, fmt.Errorf("error creating users table: %w", err)
	}

	if err := createTokensTable(db); err != nil {
		return nil, fmt.Errorf("error creating tokens table: %w", err)
	}

	if err := insertTestUser(db); err != nil {
		return nil, fmt.Errorf("error inserting test user: %w", err)
	}

	return db, nil
}
