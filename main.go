package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
)

// Global variables
var (
	todos = Todos{
		Todo{Id: 1, Title: "Write presentation", Completed: false},
		Todo{Id: 2, Title: "Host meetup", Completed: false},
	}
)

func main() {
	// Router, from gorilla/mux
	router := NewRouter()

	// Serve static files in public directory
	middleware := negroni.New(negroni.NewRecovery(),
		negroni.NewStatic(http.Dir("public")),
		// Even more middleware
		negroni.NewLogger())
	middleware.UseHandler(router)
	// Server started
	middleware.Run(":9090")
}
