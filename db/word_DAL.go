package db

import (
	"database/sql"
)

type DBBase struct {
}

func QueryAllWords() []any {
	sqlStr := "SELECT word from word"
	return Query(sqlStr, func(rows *sql.Rows) any {
		var result string
		rows.Scan(&result)
		return result
	})
}
