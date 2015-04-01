package main

import (
	"database/sql"

	// Imported a package solely for side-effects
	_ "github.com/mattn/go-sqlite3"
)

var (
	CREATE_TABLE     = "CREATE TABLE IF NOT EXISTS todos(id int, title text, completed bool)"
	INSERT_TODO      = "INSERT INTO todos (title, completed) VALUES (?, ?)"
	UPDATE_TODO      = "UPDATE todos SET title = ?, completed = ? where id = ?"
	DELETE_TODO      = "DELETE FROM todos where id = ?"
	SELECT_ALL_TODOS = "SELECT * FROM todos"
	SELECT_TODO      = "SELECT title, completed FROM todos WHERE id = ?"
	MAX_ID           = "SELECT id FROM todos ORDER BY ID DESC LIMIT 1"
	todos            Todos
	db               *sql.DB
)

// Give us some seed data
// The function init() will be automatically executed when a
// package is loaded, one go program could contains one or more init()
// function, they are executed before the actual program begins.
// http://golang.org/ref/spec#Package_initialization
func init() {
	// Create DB table if not exist
	db = GetDBConn(db)
	_, err := db.Exec(CREATE_TABLE)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func FindAllTodos() Todos {
	db = GetDBConn(db)
	defer db.Close()

	var result Todos
	rows, err := db.Query(SELECT_ALL_TODOS)
	checkErr(err)

	for rows.Next() {
		var (
			id        int
			title     string
			completed bool
			todo      Todo
		)

		err = rows.Scan(&id, &title, &completed)
		checkErr(err)
		todo = Todo{id, title, completed}
		result = append(result, todo)
	}
	return result
}

func FindTodo(id int) Todo {
	db = GetDBConn(db)
	defer db.Close()
	stmt, err := db.Prepare(SELECT_TODO)
	checkErr(err)
	defer stmt.Close()

	var (
		title     string
		completed bool
	)

	err = stmt.QueryRow(id).Scan(&title, &completed)
	checkErr(err)
	// return empty Todo if not found
	return Todo{id, title, completed}
}

func CreateTodo(t Todo) Todo {
	// No transaction at all
	db := GetDBConn(db)
	defer db.Close()
	stmt, err := db.Prepare(INSERT_TODO)
	defer stmt.Close()
	checkErr(err)

	res, err := stmt.Exec(t.Title, t.Completed)
	checkErr(err)

	lastId, err := res.LastInsertId()
	checkErr(err)
	// Convert int64 to int, may lost precision
	t.Id = int(lastId)
	return t
}

func UpdateTodo(id int, todo Todo) {
	db := GetDBConn(db)
	defer db.Close()

	stmt, err := db.Prepare(UPDATE_TODO)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(todo.Title, todo.Completed, id)
	checkErr(err)
}

func DestroyTodo(id int) {
	db = GetDBConn(db)
	defer db.Close()
	_, err := db.Exec(DELETE_TODO, id)
	checkErr(err)
}

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "todos.sqlite3")
	checkErr(err)

	return db
}

func GetDBConn(db *sql.DB) *sql.DB {
	if db == nil {
		return NewDB()
	} else {
		return db
	}
}
