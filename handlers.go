package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
	"net/http"
)

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
