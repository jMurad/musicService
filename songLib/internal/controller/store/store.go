package store

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jMurad/musicService/songLib/internal/model"
	"github.com/jMurad/musicService/songLib/pkg/postgres"
)

// type Filter [2]string
// type Filters []Filter

type Filters []struct {
	Field     string
	Operators string
	Value     any
}

type Pagination struct {
	Limit  int
	Offset int
}

const (
	Equal        = "="
	NotEqual     = "<>"
	GreaterThan  = ">"
	LessThan     = "<"
	GreaterEqual = ">="
	LessEqual    = "<="
	Like         = "LIKE"
)

type SongStore struct {
	db *postgres.Postgres
}

func NewSongStore(db *postgres.Postgres) *SongStore {
	return &SongStore{
		db: db,
	}
}

func (s *SongStore) AddSong(ctx context.Context, new *model.Song) error {
	_, err := s.db.DB().ExecContext(
		ctx,
		"INSERT INTO songs (group_name, song_name, release_date, lyrics, link) VALUES($1, $2, $3, $4, $5)",
		new.GroupName,
		new.SongName,
		new.ReleaseDate,
		new.Lyrics,
		new.Link,
	)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint`) {
			return ErrSongExists
		}
		return err
	}
	return nil
}

func (s *SongStore) EditSong(ctx context.Context, old, new *model.Song) error {
	queryUpd := "UPDATE songs SET " +
		"group_name = COALESCE(NULLIF($1, ''), group_name), " +
		"song_name = COALESCE(NULLIF($2, ''), song_name), " +
		"release_date = COALESCE(NULLIF($3, '0001-01-01'::date)::date, release_date), " +
		"lyrics = COALESCE(NULLIF($4, ''), lyrics), " +
		"link = COALESCE(NULLIF($5, ''), link) " +
		"WHERE group_name = $6 AND song_name = $7 " +
		"RETURNING song_id"

	vals := []any{
		new.GroupName,
		new.SongName,
		new.ReleaseDate,
		new.Lyrics,
		new.Link,
		old.GroupName,
		old.SongName,
	}

	err := s.db.DB().QueryRowContext(ctx, queryUpd, vals...).Scan(&new.ID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return ErrSongNotFound
		}
		return err
	}

	return nil
}

func (s *SongStore) DeleteSong(ctx context.Context, del *model.Song) error {
	err := s.db.DB().QueryRowContext(
		ctx,
		"DELETE FROM songs WHERE group_name=$1 AND song_name=$2",
		del.GroupName,
		del.SongName,
	).Scan(&del.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *SongStore) GetLyrics(ctx context.Context, song *model.Song) error {
	err := s.db.DB().QueryRowContext(
		ctx,
		"SELECT lyrics FROM songs WHERE group_name=$1 AND song_name=$2",
		song.GroupName,
		song.SongName,
	).Scan(&song.Lyrics)
	if err != nil {
		return err
	}

	return nil
}

func (s *SongStore) GetSongs(ctx context.Context, filters Filters, pag Pagination) ([]*model.Song, error) {
	songs := []*model.Song{}

	if cols, vals := columnsForFilter(filters); vals != nil {
		queryFlt := fmt.Sprintf(
			"SELECT * FROM songs WHERE %s LIMIT %d OFFSET %d",
			cols,
			pag.Limit,
			pag.Offset,
		)

		rows, err := s.db.DB().QueryContext(ctx, queryFlt, vals...)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var song model.Song
			if err := rows.Scan(&song.ID, &song.GroupName, &song.SongName, &song.ReleaseDate, &song.Lyrics, &song.Link); err != nil {
				return nil, err
			}
			songs = append(songs, &song)
		}
		if err = rows.Err(); err != nil {
			return songs, err
		}

	} else {
		return nil, errors.New("no filters")
	}

	return songs, nil
}
