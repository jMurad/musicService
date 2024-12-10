package store_test

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jMurad/musicService/songLib/internal/config"
	"github.com/jMurad/musicService/songLib/internal/controller/store"
	"github.com/jMurad/musicService/songLib/internal/model"
	"github.com/jMurad/musicService/songLib/pkg/postgres"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"

	"github.com/brianvoe/gofakeit/v7"
)

const (
	UniqueConstraint = pq.ErrorCode("23505")
)

func IsErrorCode(err error, errcode pq.ErrorCode) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == errcode
	}
	return false
}

var st *store.SongStore

func TestMain(m *testing.M) {
	os.Setenv("CONFIG_PATH", "/Users/murad/goProjects/projects/musicService/songLib/config/config.yaml")
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalln(err)
		return
	}

	tlog := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	db, err := postgres.New(cfg.StoragePath, tlog)
	if err != nil {
		log.Fatalln(err)
		return
	}

	st = store.NewSongStore(db)

	m.Run()
}

func newRandomSong() *model.Song {
	rDate := gofakeit.Date()

	return &model.Song{
		GroupName:   gofakeit.Product().Name,
		SongName:    gofakeit.GlobalFaker.Name(),
		ReleaseDate: time.Date(rDate.Year(), rDate.Month(), rDate.Day(), 0, 0, 0, 0, rDate.Location()),
		Lyrics:      lyricsGenerate(),
	}
}

func lyricsGenerate() string {
	rand.Seed(uint64(time.Now().UnixNano()))

	alpha := "abcdefghijklmnopqrstuvwxyz"

	var lyrics strings.Builder
	k := len(alpha)

	countWords := 4
	countLines := 4
	countCouplets := 4

	for couplets := 0; couplets < countCouplets; couplets++ {
		for lines := 0; lines < countLines; lines++ {
			for words := 0; words < countWords; words++ {
				for characters := 0; characters < 2+rand.Intn(8); characters++ {
					c := alpha[rand.Intn(k)]
					lyrics.WriteByte(c)
				}
				lyrics.WriteString(" ")
			}
			lyrics.WriteString("\n")
		}
		lyrics.WriteString("\n")
	}
	return lyrics.String()
}

func CreateRandomSong(t *testing.T) *model.Song {
	newSong := newRandomSong()

	song, err := st.AddSong(context.Background(), newSong)
	assert.NoError(t, err)
	assert.NotNil(t, song)
	assert.Equal(t, newSong.GroupName, song.GroupName)
	assert.Equal(t, newSong.SongName, song.SongName)
	assert.Equal(t, newSong.ReleaseDate, song.ReleaseDate.UTC())
	assert.Equal(t, newSong.Lyrics, song.Lyrics)
	assert.Equal(t, newSong.Link, song.Link)

	return song
}

func TestAddSong(t *testing.T) {
	testCases := []struct {
		name       string
		testScript func()
	}{
		{"create random song", func() {
			CreateRandomSong(t)
		}},
		{"unique GroupName and SongName error", func() {
			newSong := CreateRandomSong(t)

			violationSong, err := st.AddSong(context.Background(), newSong)
			assert.Equal(t, true, IsErrorCode(err, UniqueConstraint))
			assert.Nil(t, violationSong)

		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.testScript()
		})
	}
}

func TestEditSong(t *testing.T) {
	newSong := newRandomSong()
	song, _ := st.AddSong(context.Background(), newSong)

	oldSong := &model.Song{}
	oldSong.GroupName = newSong.GroupName
	oldSong.SongName = newSong.SongName

	song, err := st.EditSong(context.Background(), oldSong, newSong)
	assert.Error(t, err)
	assert.NotNil(t, song)
	assert.Equal(t, newSong.GroupName, song.GroupName)
	assert.Equal(t, newSong.SongName, song.SongName)
	assert.Equal(t, newSong.ReleaseDate, song.ReleaseDate)
	assert.Equal(t, newSong.Lyrics, song.Lyrics)
	assert.Equal(t, newSong.Link, song.Link)

}

func TestDeleteSong(t *testing.T) {
	song := model.Song{}
	err := st.DeleteSong(context.Background(), &song)
	if err != nil {
		t.Error(err)
	}
}

func TestGetLyrics(t *testing.T) {
	song := model.Song{}
	lyrics := song
	lyrics.Lyrics = ""
	err := st.GetLyrics(context.Background(), &lyrics)
	if err != nil {
		t.Error(err)
	}
	if song.Lyrics != lyrics.Lyrics {
		t.Error("Lyrics:", lyrics.Lyrics)
	}
}

func TestGetSongs(t *testing.T) {
	var filters = store.Filters{
		{
			Field:     "song_name",
			Operators: store.Like,
			Value:     "%n%",
		},
		{
			Field:     "release_date",
			Operators: store.LessEqual,
			Value:     toDate("30.11.2024"),
		},
		{
			Field:     "release_date",
			Operators: store.GreaterEqual,
			Value:     toDate("25.11.2024"),
		},
		{
			Field:     "lyrics",
			Operators: store.Like,
			Value:     "%ly1%",
		},
	}

	pag := store.Pagination{
		Limit:  10,
		Offset: 0,
	}

	songs, err := st.GetSongs(context.Background(), filters, pag)
	if err != nil {
		t.Error(err)
	}

	for _, v := range songs {
		fmt.Println(v)
	}
	t.Error("---")
}

func toDate(strDate string) time.Time {
	dt, err := time.Parse("02.01.2006", strDate)
	if err != nil {
		return time.Now()
	}
	return dt
}
