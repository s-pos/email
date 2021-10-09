package repository

import "github.com/jmoiron/sqlx"

type repo struct {
	db *sqlx.DB
}

type Repo interface{}

func New(db *sqlx.DB) Repo {
	return &repo{
		db: db,
	}
}
