package pgstore_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/jMurad/musicService/songLib/internal/config"
	"github.com/jMurad/musicService/songLib/internal/store/pgstore"
)

func TestNewStore(t *testing.T) {
	cfg := config.MustLoad()

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	_, err := pgstore.NewStore(cfg.StoragePath, log)
	if err != nil {
		t.Error(err)
	}

}
