package main

import (
	"database/sql"

	// Imported a package solely for side-effects
	_ "github.com/mattn/go-sqlite3"
)

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "todos.sqlite3")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("create table if not exists books(id int, content text, complete bool)")
	if err != nil {
		panic(err)
	}

	return db
}
