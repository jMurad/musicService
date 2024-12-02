package pgstore

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

type store struct {
	db  *sql.DB
	log *slog.Logger
}

func NewStore(storePath string, log *slog.Logger) (*store, error) {
	db, err := sql.Open("postgres", storePath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &store{
		db:  db,
		log: log,
	}, nil
}
