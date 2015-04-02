package main

import (
	"database/sql"

	// Imported a package solely for side-effects
	_ "github.com/mattn/go-sqlite3"
)

var (
	CREATE_TABLE     = "CREATE TABLE IF NOT EXISTS todos(id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, completed BOOL)"
	INSERT_TODO      = "INSERT INTO todos (title, completed) VALUES (?, ?)"
	UPDATE_TODO      = "UPDATE todos SET title = ?, completed = ? where id = ?"
	DELETE_TODO_ITEM = "DELETE FROM todos where id = ?"
	DELETE_TODOs     = "DELETE FROM todos where id in ( ? )"
	SELECT_ALL_TODOS = "SELECT * FROM todos"
	SELECT_TODO      = "SELECT title, completed FROM todos WHERE id = ?"
	MAX_ID           = "SELECT id FROM todos ORDER BY ID DESC LIMIT 1"
	db               *sql.DB
)

// Give us some seed data
// The function init() will be automatically executed when a
// package is loaded, one go program could contains one or more init()
// function, they are executed before the actual program begins.
// http://golang.org/ref/spec#Package_initialization
func init() {
	// As this article says: http://go-database-sql.org/accessing.html
	// the sql.DB object is designed to be long-lived. Don’t Open() and Close() databases frequently.
	// Instead, create one sql.DB object for each distinct datastore you need to access,
	// and keep it until the program is done accessing that datastore. Pass it around as needed,
	// or make it available somehow globally, but keep it open.
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

func UpdateTodo(todo Todo) {
	stmt, err := db.Prepare(UPDATE_TODO)
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(todo.Title, todo.Completed, todo.Id)
	checkErr(err)
}

func DestroyTodoItem(id int) {
	_, err := db.Exec(DELETE_TODO_ITEM, id)
	checkErr(err)
}

func DestroyTodos(ids string) {
	_, err := db.Exec(DELETE_TODOs, ids)
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
