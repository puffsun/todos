package main

import (
	"github.com/russross/blackfriday"
	"net/http"
)

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	// Serve static files in public directory
	http.Handle("/", http.FileServer(http.Dir("public")))

	println("Go web server up and running...")
	http.ListenAndServe(":9090", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
