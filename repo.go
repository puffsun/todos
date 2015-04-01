package main

import (
	"database/sql"
	"fmt"

	// Imported a package solely for side-effects
	_ "github.com/mattn/go-sqlite3"
)

var (
	CREATE_TABLE = "create table if not exists books(id int, content text, complete bool)"
)

var currentId int

var todos Todos

// Give us some seed data
func init() {
	RepoCreateTodo(Todo{Title: "Write presentation"})
	RepoCreateTodo(Todo{Title: "Host meetup"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "todos.sqlite3")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(CREATE_TABLE)
	if err != nil {
		panic(err)
	}

	return db
}
