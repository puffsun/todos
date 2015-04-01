package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
	"net/http"
)

func GenerateMarkdown(resp http.ResponseWriter, req *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(req.FormValue("body")))
	resp.Write(markdown)
}

func TodosIndexHandler(resp http.ResponseWriter, req *http.Request) {
	// net/http server will always set accurate content-type and status code
	// itself  but we set it explicitly here to make it clear.
	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(resp).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoItemGetHandler(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	todoId := vars["todoId"]
	fmt.Fprintln(resp, "Todo item get handler: ", todoId)
}

func TodosPostHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Post Handler")
}

func TodosPutHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Put Handler")
}

func TodosDeleteHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Delete Handler")
}

func ApiHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "API Handler")
}
