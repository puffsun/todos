package main

import (
	"github.com/codegangsta/negroni"
	"net/http"
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
