package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DBconect() {
	var db, err = sql.Open("mysql", "root:root@(127.0.0.1:8889)/findme?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
