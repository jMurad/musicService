package store

import "errors"

var (
	ErrSongNotFound = errors.New("song not found")
	ErrSongExists   = errors.New("song exists")
)
