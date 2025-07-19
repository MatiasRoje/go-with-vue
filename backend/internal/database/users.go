package database

import (
	"database/sql"

	"github.com/MatiasRoje/go-with-vue/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UsersModel struct {
	DB *sql.DB
}

func (m *UsersModel) GetAll() ([]models.User, error) {
	rows, err := m.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (m *UsersModel) GetByEmail(email string) (models.User, error) {
	row := m.DB.QueryRow("SELECT * FROM users WHERE email = $1", email)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}
	return user, nil
}

func (m *UsersModel) GetByID(id int) (models.User, error) {
	row := m.DB.QueryRow("SELECT * FROM users WHERE id = $1", id)
	var user models.User
	if err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}
	return user, nil
}

func (m *UsersModel) Insert(user models.User) error {
	query := `
		INSERT INTO users (email, first_name, last_name, password)
		VALUES ($1, $2, $3, $4)
	`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if _, err := m.DB.Exec(query, user.Email, user.FirstName, user.LastName, hashedPassword); err != nil {
		return err
	}
	return nil
}
