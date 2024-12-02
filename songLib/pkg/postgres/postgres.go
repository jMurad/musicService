package postgres

import (
	"context"
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

type DB interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Postgres struct {
	db *sql.DB
}

func New(storePath string, log *slog.Logger) (*Postgres, error) {
	db, err := sql.Open("postgres", storePath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) Close() {
	p.db.Close()
}

func (p *Postgres) DB() DB {
	return p.db
}
