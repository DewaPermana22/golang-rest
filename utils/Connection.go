package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connection() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/rest-go")
	if err != nil {
		panic(err)
	}

	DB = db
}

// // See "Important settings" section.
// db.SetConnMaxLifetime(time.Minute * 3)
// db.SetMaxOpenConns(10)
// db.SetMaxIdleConns(10)
