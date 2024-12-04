package store_test

import (
	"context"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/jMurad/musicService/songLib/internal/config"
	"github.com/jMurad/musicService/songLib/internal/controller/store"
	"github.com/jMurad/musicService/songLib/internal/model"
	"github.com/jMurad/musicService/songLib/pkg/postgres"
)

var pg *postgres.Postgres
var cfg *config.Config
var log *slog.Logger
var st *store.SongStore
var song model.Song

func TestMain(m *testing.M) {
	os.Setenv("CONFIG_PATH", "/Users/murad/goProjects/projects/musicService/songLib/config/config.yaml")
	cfg, _ = config.MustLoad()

	log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	pg, _ = postgres.New(cfg.StoragePath, log)
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
	newSong := song

	err := st.AddSong(context.Background(), &newSong)
	if err != nil {
		t.Error(err)
	}
}

func TestEditSong(t *testing.T) {
	editSong := song
	editSong.GroupName = "Your Group"
	editSong.SongName = "Your Song"
	editSong.Link = "https://url.your/editsong"

	err := st.EditSong(context.Background(), &song, &editSong)
	if err != nil {
		t.Error(err)
	}
}
