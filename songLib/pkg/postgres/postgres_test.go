package pgstore_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/jMurad/musicService/songLib/internal/config"
	"github.com/jMurad/musicService/songLib/internal/store/pgstore"
)

var cfg *config.Config

func TestMain(m *testing.M) {
	os.Setenv("CONFIG_PATH", "/Users/murad/goProjects/projects/musicService/songLib/config/config.yaml")
	cfg = config.MustLoad()

	os.Exit(m.Run())
}

func TestNewStore(t *testing.T) {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	_, err := pgstore.NewStore(cfg.StoragePath, log)
	if err != nil {
		t.Error(err)
	}

}

type Store interface {
	AddSong(ctx context.Context, song pgstore.songs) error
}

func TestAddSong(t *testing.T) {
	cfg := config.MustLoad()

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	_, err := pgstore.NewStore(cfg.StoragePath, log)
	if err != nil {
		t.Error(err)
	}

}
