package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// Book is the definition of a single book
type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	AuthorID        int       `json:"author_id"`
	PublicationYear int       `json:"publication_year"`
	Slug            string    `json:"slug"`
	Author          Author    `json:"author"`
	Description     string    `json:"description"`
	Genres          []*Genre  `json:"genres"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Author is the definition of a single author
type Author struct {
	ID         int       `json:"id"`
	AuthorName string    `json:"author_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Genre is the definition of a single genre type
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
