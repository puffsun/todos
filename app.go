package main

import (
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	// Imported a package solely for side-effects
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/russross/blackfriday"
	"net/http"
)

// Global variables
var (
	TODOS_PATH = "/api/todos"
	TODO_PATH  = "/api/todos/{todoId}"
	todos      = Todos{
		Todo{Id: 1, Title: "Write presentation", Completed: false},
		Todo{Id: 2, Title: "Host meetup", Completed: false},
	}
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Todos []Todo

func main() {

	db := NewDB()
	println(db)
	// Router, from gorilla/mux
	router := mux.NewRouter().StrictSlash(false)

	// Markdown generator, from russross/blackfriday
	markdown := router.Path("/markdown").Subrouter()
	markdown.Methods("POST").HandlerFunc(GenerateMarkdown)

	api := router.Path("/api").Subrouter()
	api.Methods("GET").HandlerFunc(ApiHandler)

	todos := router.Path(TODOS_PATH).Subrouter()
	todos.Methods("GET").HandlerFunc(TodosIndexHandler)

	todoItem := router.Path(TODO_PATH).Subrouter()
	todoItem.Methods("GET").HandlerFunc(TodoItemGetHandler)
	todoItem.Methods("POST").HandlerFunc(TodosPostHandler)
	todoItem.Methods("DELETE").HandlerFunc(TodosDeleteHandler)
	todoItem.Methods("PUT").HandlerFunc(TodosPutHandler)

	// Serve static files in public directory
	middleware := negroni.New(negroni.NewRecovery(),
		negroni.NewStatic(http.Dir("public")),
		// Even more middleware
		negroni.NewLogger())

	middleware.UseHandler(router)

	// Server started
	middleware.Run(":9090")
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

func TodosIndexHandler(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(todos)
}

func TodoItemGetHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(rw, "Todo item get handler: ", todoId)
}

func TodosPostHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Post Handler")
}

func TodosPutHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Put Handler")
}

func TodosDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Delete Handler")
}

func ApiHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "API Handler")
}

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
