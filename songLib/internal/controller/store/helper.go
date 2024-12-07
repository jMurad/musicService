package store

import (
	"fmt"
	"strings"
)

func columnsForFilter(filters Filters) (string, []any) {
	columns := []string{}
	vals := []any{}
	for i, filter := range filters {
		columns = append(
			columns,
			fmt.Sprintf("%s %s $%d", filter.Field, filter.Operators, i+1),
		)
		vals = append(vals, filter.Value)
	}

	return strings.Join(columns, " AND "), vals
}
