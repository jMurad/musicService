package store_test

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/jMurad/musicService/songLib/internal/config"
	"github.com/jMurad/musicService/songLib/internal/controller/store"
	"github.com/jMurad/musicService/songLib/internal/model"
	"github.com/jMurad/musicService/songLib/pkg/postgres"
)

var st *store.SongStore
var song model.Song

func TestMain(m *testing.M) {
	os.Setenv("CONFIG_PATH", "/Users/murad/goProjects/projects/musicService/songLib/config/config.yaml")
	cfg, _ := config.MustLoad()

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	pg, _ := postgres.New(cfg.StoragePath, log)
	st = store.NewSongStore(pg)

	song = model.Song{
		GroupName:   "My Group",
		SongName:    "My Song",
		ReleaseDate: time.Now(),
		Lyrics:      "la la la",
		Link:        "https://url.my/song",
	}

	os.Exit(m.Run())
}

func TestAddSong(t *testing.T) {
	newSong := model.Song{}
	// newSong.GroupName = "gn1"
	// newSong.SongName = "sn1"
	// newSong.Lyrics = "ly ly ly"
	// newSong.Link = "new"

	err := st.AddSong(context.Background(), &newSong)
	if err != nil {
		t.Error(err)
	}
}

func TestEditSong(t *testing.T) {

	old := model.Song{
		GroupName: "gn1",
		SongName:  "sn1",
	}
	editSong := model.Song{}
	editSong.Link = "EDIT"
	// editSong.ReleaseDate = toDate("04.12.2024")

	err := st.EditSong(context.Background(), &old, &editSong)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteSong(t *testing.T) {
	err := st.DeleteSong(context.Background(), &song)
	if err != nil {
		t.Error(err)
	}
}

func TestGetLyrics(t *testing.T) {
	lyrics := song
	lyrics.Lyrics = ""
	err := st.GetLyrics(context.Background(), &lyrics)
	if err != nil {
		t.Error(err)
	}
	if song.Lyrics != lyrics.Lyrics {
		t.Error("Lyrics:", lyrics.Lyrics)
	}
}

func TestGetSongs(t *testing.T) {
	var filters = store.Filters{
		{
			Field:     "song_name",
			Operators: store.Like,
			Value:     "%n%",
		},
		{
			Field:     "release_date",
			Operators: store.LessEqual,
			Value:     toDate("30.11.2024"),
		},
		{
			Field:     "release_date",
			Operators: store.GreaterEqual,
			Value:     toDate("25.11.2024"),
		},
		{
			Field:     "lyrics",
			Operators: store.Like,
			Value:     "%ly1%",
		},
	}

	pag := store.Pagination{
		Limit:  10,
		Offset: 0,
	}

	songs, err := st.GetSongs(context.Background(), filters, pag)
	if err != nil {
		t.Error(err)
	}

	for _, v := range songs {
		fmt.Println(v)
	}
	t.Error("---")
}

func toDate(strDate string) time.Time {
	dt, err := time.Parse("02.01.2006", strDate)
	if err != nil {
		return time.Now()
	}
	return dt
}
