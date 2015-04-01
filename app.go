package main

import (
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	// Imported a package solely for side-effects
	_ "github.com/mattn/go-sqlite3"
	"github.com/russross/blackfriday"
	"net/http"
)

// Global variables
var (
	TODOS_PATH = "/api/todos"
	todos      = Todos{
		Todo{Id: "1", Title: "Write presentation", Completed: false},
		Todo{Id: "2", Title: "Host meetup", Completed: false},
	}
)

type Todo struct {
	Id        string `json:"id"`
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
	todos.Methods("GET").HandlerFunc(TodosGetHandler)
	todos.Methods("POST").HandlerFunc(TodosPostHandler)
	todos.Methods("DELETE").HandlerFunc(TodosDeleteHandler)
	todos.Methods("PUT").HandlerFunc(TodosPutHandler)

	// Serve static files in public directory
	//http.Handle("/", http.FileServer(http.Dir("public")))
	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	middleware := negroni.New(negroni.NewRecovery(),
		negroni.NewStatic(http.Dir("public")),
		// Even more middleware
		negroni.NewLogger())

	middleware.UseHandler(router)

	// Server started
	//http.ListenAndServe(":9090", router)
	middleware.Run(":9090")
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

func TodosGetHandler(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(todos)
}

func TodosPostHandler(rw http.ResponseWriter, r *http.Request) {
	println("Post request accepted.")

}

func TodosPutHandler(rw http.ResponseWriter, r *http.Request) {
	println("Put request accepted.")
}

func TodosDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	println("Delete request accepted.")
}

func ApiHandler(rw http.ResponseWriter, r *http.Request) {
	println("API request accepted.")
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
