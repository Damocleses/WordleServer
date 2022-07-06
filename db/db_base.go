package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type execFunc func(db *sql.DB) []any
type readDataFunc func(rows *sql.Rows) any

func Query(sqlStr string, readFunc readDataFunc) []any {
	return execSQL(func(db *sql.DB) []any {
		rows, err := db.Query(sqlStr)
		checkErr(err)
		var result []any
		for rows.Next() {
			data := readFunc(rows)
			result = append(result, data)
		}
		return result
	})
}

func execSQL(action execFunc) []any {
	db, err := sql.Open("mysql", "root:Qzj!1234@tcp(152.136.215.195:3306)/manager_platform")
	checkErr(err)
	result := action(db)
	defer db.Close()
	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
