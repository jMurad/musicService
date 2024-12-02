package v1

import (
	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	mwLog "github.com/jMurad/musicService/songLib/pkg/logger/middleware"
)

type Usecase interface {
	Song()
}

func NewRouter(router *chi.Mux, l *slog.Logger) {
	router.Use(middleware.RequestID)
	router.Use(mwLog.New(l))
	router.Use(middleware.Recoverer)

	router.Route("/song", func(r chi.Router) {
		r.Get("/", Song(log, svc))
		r.Get("/lyrics", Lyrics(log, svc))
		r.Post("/", Add(log, svc))
		r.Put("/", Edit(log, svc))
		r.Delete("/", Delete(log, svc))
	})
}
