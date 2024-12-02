package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/jMurad/musicService/songLib/internal/config"
	v1 "github.com/jMurad/musicService/songLib/internal/controller/api/v1"
	"github.com/jMurad/musicService/songLib/pkg/httpserver"
	"github.com/jMurad/musicService/songLib/pkg/logger"
)

type App struct {
	log    *slog.Logger
	server *http.Server
}

func Run(cfg *config.Config) {
	slogger := logger.SetupLogger(cfg.Env)

	// HTTP Server
	handler := chi.NewRouter()
	v1.NewRouter(handler, slogger)
	httpServer := httpserver.New(cfg, handler)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error

	select {
	case s := <-interrupt:
		slogger.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		slogger.Error(fmt.Sprintf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		slogger.Error(fmt.Sprintf("app - Run - httpServer.Shutdown: %w", err))
	}
}
