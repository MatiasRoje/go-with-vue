package database

import "database/sql"

type Models struct {
	DBUsers DBUsers
	DBBooks DBBooks
}

type DBUsers struct {
	DB *sql.DB
}

type DBBooks struct {
	DB *sql.DB
}

func NewDBModels(db *sql.DB) *Models {
	return &Models{
		DBUsers: DBUsers{DB: db},
		DBBooks: DBBooks{DB: db},
	}
}
