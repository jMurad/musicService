package postgres

import (
	"context"
	"database/sql"
	"log/slog"

	"time"

	_ "github.com/lib/pq"
)

type songs struct {
	GroupName   string    `db:"group_name,omitempty"`
	SongName    string    `db:"song_name,omitempty"`
	ReleaseDate time.Time `db:"release_date,omitempty"`
	Lyrics      string    `db:"lyrics,omitempty"`
	Link        string    `db:"link,omitempty"`
}

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

func (s store) AddSong(ctx context.Context, song songs) error {
	_, err := s.db.ExecContext(ctx, "insert into songs (group_name, song_name, release_date, lyrics, link) values(?,?,?,?,?)",
		song.GroupName,
		song.SongName,
		song.ReleaseDate,
		song.Lyrics,
		song.Link,
	)
	if err != nil {
		return err
	}

	return nil
}
