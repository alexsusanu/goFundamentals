package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "static/contact.html")
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w)
	case "/contact":
		contactHandler(w, r)
	default:
		//errorHandling(w, http.StatusNotFound)
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting on :3000")
	http.ListenAndServe(":3000", router)
}
