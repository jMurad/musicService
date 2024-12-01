package app

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jMurad/musicService/songLib/internal/service"
	"github.com/jMurad/musicService/songLib/internal/store/pgstore"
	"github.com/jMurad/musicService/songLib/pkg/middleware/logger"
)

type App struct {
	log    *slog.Logger
	server *http.Server
}

func New(
	log *slog.Logger,
	storePath string,
	address string,
	rwTimeout time.Duration,
	idleTimeout time.Duration,
) *App {
	store, err := pgstore.NewStore(storePath, log)
	log.Error("failed to open db", slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	})

	svc, err := service.NewService(store)
	log.Error("failed to init service", slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	})

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(logger.New(log))
	router.Use(middleware.Recoverer)

	router.Route("/song", func(r chi.Router) {
		r.Get("/", Song(log, svc))
		r.Get("/lyrics", Lyrics(log, svc))
		r.Post("/", Add(log, svc))
		r.Put("/", Edit(log, svc))
		r.Delete("/", Delete(log, svc))
	})

	srv := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  rwTimeout,
		WriteTimeout: rwTimeout,
		IdleTimeout:  idleTimeout,
	}

	return &App{
		log:    log,
		server: srv,
	}
}

func (a *App) Start() error {
	return nil
}
