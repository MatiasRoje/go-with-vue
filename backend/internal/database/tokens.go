package database

import (
	"database/sql"

	"github.com/MatiasRoje/go-with-vue/backend/internal/models"
)

type TokensModel struct {
	DB *sql.DB
}

func (m *TokensModel) GetAll() ([]models.Token, error) {
	rows, err := m.DB.Query("SELECT * FROM tokens")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []models.Token
	for rows.Next() {
		var token models.Token
		if err := rows.Scan(&token.ID, &token.UserID, &token.UserEmail, &token.Token, &token.TokenHash, &token.CreatedAt, &token.UpdatedAt, &token.ExpiresAt); err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}
