package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// NOTE: The logic around JWT should be moved to a separate package/service in a real application

const (
	JWT_EXPIRATION = 24 * time.Hour
)

type tokenClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func generateJWTToken(userEmail, secret string) (string, error) {
	expiresAt := time.Now().Add(JWT_EXPIRATION)
	claims := tokenClaims{
		Email: userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-with-vue",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateJWTToken(tokenString string, secret string) (*tokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*tokenClaims); ok && token.Valid {
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, fmt.Errorf("token expired")
		}
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
