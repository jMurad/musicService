package store

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jMurad/musicService/songLib/internal/model"
)

func columnsForUpdate(old, new any) (string, []any) {
	updateColumns := ""
	vals := []any{}
	counter := 0

	oldV := reflect.ValueOf(old)
	newV := reflect.ValueOf(new)
	songV := reflect.ValueOf(model.Song{}).Type()

	if songV != oldV.Type() || songV != newV.Type() {
		return "", nil
	}

	for i := 0; i < songV.NumField(); i++ {
		if !oldV.Field(i).Equal(newV.Field(i)) {
			counter++
			updateColumns += fmt.Sprintf(" %s = $%d,", strings.Split(songV.Field(i).Tag.Get("json"), ",")[0], counter)
			vals = append(vals, newV.Field(i).Interface())
		}
	}

	if len(vals) == 0 {
		return "", nil
	}

	fmt.Println("vals:", vals)

	return strings.TrimSuffix(updateColumns, ","), vals
}
