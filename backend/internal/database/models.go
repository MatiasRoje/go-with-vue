package database

import "database/sql"

type Models struct {
	Users UsersModel
}

func NewDBModels(db *sql.DB) *Models {
	return &Models{
		Users: UsersModel{DB: db},
	}
}
