package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}

var (
	TODOS_PATH = "/api/todos"
	TODO_PATH  = "/api/todos/{todoId}"
)
var routes = Routes{
	Route{
		"markdown_generator", "GET", "/markdown", GenerateMarkdown,
	},
	Route{
		"API", "GET", "/api", ApiHandler,
	},
	Route{
		"todo_index", "GET", TODOS_PATH, TodosIndexHandler,
	},
	Route{
		"todo_item_get", "GET", TODO_PATH, TodoItemGetHandler,
	},
	Route{
		"todo_item_post", "POST", TODO_PATH, TodosPostHandler,
	},
	Route{
		"todo_item_delete", "DELETE", TODO_PATH, TodosDeleteHandler,
	},
	Route{
		"todo_item_put", "PUT", TODO_PATH, TodosPutHandler,
	},
}
