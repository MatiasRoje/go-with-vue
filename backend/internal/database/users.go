package database

import (
	"context"

	"github.com/MatiasRoje/go-with-vue/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *DBUsers) GetAll() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (m *DBUsers) GetByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE email = $1", email)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return &user, err
	}
	return &user, nil
}

func (m *DBUsers) GetByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE id = $1", id)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return &user, err
	}
	return &user, nil
}

func (m *DBUsers) Insert(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	query := `
		INSERT INTO users (email, first_name, last_name, password)
		VALUES ($1, $2, $3, $4)
	`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if _, err := m.DB.ExecContext(ctx, query, user.Email, user.FirstName, user.LastName, hashedPassword); err != nil {
		return err
	}
	return nil
}
