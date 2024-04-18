package store

import (
	"context"
	"database/sql"
)

type DBOperations interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

func NewQueries(db DBOperations) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db DBOperations
}
