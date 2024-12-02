package model

import "time"

type Song struct {
	GroupName string `json:"group_name,omitempty"`
	SongName  string `json:"song_name,omitempty"`
}

type SongInfo struct {
	ReleaseDate time.Time `json:"release_date,omitempty"`
	Lyrics      string    `json:"lyrics,omitempty"`
	Link        string    `json:"link,omitempty"`
}
