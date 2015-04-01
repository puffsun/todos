package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
	"io"
	"io/ioutil"
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

func TodosCreateHandler(resp http.ResponseWriter, req *http.Request) {
	var todo Todo
	// protect against malicious attacks on your server,
	// imagine some malicious user send us a 500GB JSON
	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := req.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resp.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(resp).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	resp.Header().Set("Content-Type", "application/json; charset=UTF-8")
	resp.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(resp).Encode(t); err != nil {
		panic(err)
	}
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
