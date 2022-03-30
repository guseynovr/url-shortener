package main

import (
	// "os"
	"net/http"
	// "html/template"
)

func main() {
	// http.HandleFunc("/", mainPage)
	http.ListenAndServe(":8080", http.HandlerFunc(mainPage))
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
