package store

import (
	"testing"
	"time"
)

func TestColumnsForFilter(t *testing.T) {
	var filters = Filters{
		{
			Field:     "song_name",
			Operators: Like,
			Value:     "Your",
		},
		{
			Field:     "release_date",
			Operators: LessEqual,
			Value:     toDate("03.12.2024"),
		},
		{
			Field:     "lyrics",
			Operators: Like,
			Value:     "la",
		},
	}
	test := "song_name LIKE $1 AND release_date <= $2 AND lyrics LIKE $3"

	q, v := columnsForFilter(filters)
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
