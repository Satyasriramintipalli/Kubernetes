package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {

	var err error

	db, err = sql.Open("sqlite3", "./users.db")

	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT
	);
	`

	statement, err := db.Prepare(createTable)

	if err != nil {
		log.Fatal(err)
	}

	statement.Exec()
}