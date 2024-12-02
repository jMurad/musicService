package postgres_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/jMurad/musicService/songLib/internal/config"
	"github.com/jMurad/musicService/songLib/pkg/postgres"
)

var cfg *config.Config
var log *slog.Logger

func TestMain(m *testing.M) {
	os.Setenv("CONFIG_PATH", "/Users/murad/goProjects/projects/musicService/songLib/config/config.yaml")
	cfg, _ = config.MustLoad()

	log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	_, err := postgres.New(cfg.StoragePath, log)
	if err != nil {
		t.Error(err)
	}

}

func TestAdd(t *testing.T) {
	db, err := postgres.New(cfg.StoragePath, log)
	if err != nil {
		t.Error(err)
	}

	_, err = db.DB().ExecContext(context.Background(), "INSERT INTO songs (group_name, song_name, release_date, lyrics, link) VALUES($1, $2, $3, $4, $5)",
		"groupName",
		"songNmae",
		"12.12.2024",
		"la la la la la",
		"http://link.com",
	)
	if err != nil {
		t.Error(err)
	}
}
