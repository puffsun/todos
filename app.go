package main

import (
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
	"net/http"
)

// Global variables
var (
	TODOS_PATH = "/api/todos"
)

func main() {

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
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// Server started
	println("Go web server up and running...")
	http.ListenAndServe(":9090", router)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}

func TodosGetHandler(rw http.ResponseWriter, r *http.Request) {
	println("Get request accepted.")
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
