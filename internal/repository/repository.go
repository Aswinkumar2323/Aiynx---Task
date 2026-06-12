package repository

import (
	"database/sql"
	"sqlc.dev/app/db/sqlc"
)

type UserRepository interface {
	Queries() *sqlc.Queries
}

type userRepository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db:      db,
		queries: sqlc.New(db),
	}
}

func (r *userRepository) Queries() *sqlc.Queries {
	return r.queries
}
