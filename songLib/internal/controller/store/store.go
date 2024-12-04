package store

import (
	"context"
	"fmt"

	"github.com/jMurad/musicService/songLib/internal/model"
	"github.com/jMurad/musicService/songLib/pkg/postgres"
)

type SongStore struct {
	db *postgres.Postgres
}

func NewSongStore(db *postgres.Postgres) *SongStore {
	return &SongStore{
		db: db,
	}
}

func (s *SongStore) AddSong(ctx context.Context, song *model.Song) error {
	_, err := s.db.DB().ExecContext(
		context.Background(),
		"INSERT INTO songs (group_name, song_name, release_date, lyrics, link) VALUES($1, $2, $3, $4, $5)",
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

func (s *SongStore) EditSong(ctx context.Context, old, new *model.Song) error {
	err := s.db.DB().QueryRowContext(
		ctx,
		"SELECT song_id FROM songs WHERE group_name=$1 AND song_name=$2",
		old.GroupName,
		old.SongName,
	).Scan(
		&old.ID,
	)
	if err != nil {
		return err
	}

	new.ID = old.ID

	if columns, vals := columnsForUpdate(*old, *new); vals != nil {
		vals = append(vals, old.ID)
		queryUpd := fmt.Sprintf("UPDATE songs SET%s WHERE song_id = $%d", columns, len(vals))
		_, err = s.db.DB().ExecContext(context.Background(), queryUpd, vals...)
		if err != nil {
			return err
		}
	}
	return nil
}
