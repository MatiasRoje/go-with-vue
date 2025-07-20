package main

import "github.com/MatiasRoje/go-with-vue/backend/internal/models"

type frontendUser struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

func userToFrontendUser(user models.User) frontendUser {
	return frontendUser{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
