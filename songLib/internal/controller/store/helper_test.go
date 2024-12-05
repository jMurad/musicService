package store

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jMurad/musicService/songLib/internal/model"
)

func TestColumnsForUpdate(t *testing.T) {
	song := model.Song{}
	song.SongName = "Your"
	song.Lyrics = "la"
	song.ReleaseDate = toDate("03.12.2024")
	test := fmt.Sprintf("%s = $1, %s = $2, %s = $3", "song_name", "release_date", "lyrics")

	q, v := columnsForUpdate(&model.Song{}, &song)
	if q != test {
		t.Errorf("\nq:\t\t|%s| with v:{%v}\ntest:\t|%s|\n doesn't match", q, v, test)

	}
}
func TestColumnsForFilter(t *testing.T) {
	filter := model.Song{}
	filter.SongName = "Your"
	filter.Lyrics = "la"
	filter.ReleaseDate = toDate("03.12.2024")

	var oper Operators = Operators{}
	oper = append(oper, Like, LessEqual, GreaterThan)

	test := fmt.Sprintf("%s {op} $1 AND %s {op} $2 AND %s {op} $3", "song_name", "release_date", "lyrics")
	for _, rpl := range oper {
		test = strings.Replace(test, "{op}", rpl, 1)
	}

	q, v := columnsForFilter(oper, &filter)
	if q != test {
		t.Errorf("\nq:\t\t|%s| with v:{%v}\ntest:\t|%s|\n doesn't match", q, v, test)

	}
}

func toDate(strDate string) time.Time {
	dt, err := time.Parse("02.01.2006", strDate)
	if err != nil {
		return time.Now()
	}
	return dt
}
