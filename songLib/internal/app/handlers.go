package app

import (
	"log/slog"
	"net/http"
)

type Service interface {
}

func Song(log *slog.Logger, svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Lyrics(log *slog.Logger, svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Add(log *slog.Logger, svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Edit(log *slog.Logger, svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Delete(log *slog.Logger, svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
