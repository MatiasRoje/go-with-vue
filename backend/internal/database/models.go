package database

import "database/sql"

type Models struct {
	Users  UsersModel
	Tokens TokensModel
}

func NewDBModels(db *sql.DB) *Models {
	return &Models{
		Users:  UsersModel{DB: db},
		Tokens: TokensModel{DB: db},
	}
}
