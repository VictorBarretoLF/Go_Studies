package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, please send an email to <a href=\"mailto:victorbarretolins@gmail.com\">victorbarretolins@gmail.com</a>.</p>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// w.WriteHeader(http.StatusNotFound)
	// fmt.Fprint(w, "<h1>Page Not Found</h1><p>The page you are looking for does not exist.</p>")
	http.Error(w, "Page not found", http.StatusNotFound)
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
		case "/":
			homeHandler(w, r)
		case "/contact":
			contactHandler(w, r)
		default:
			notFoundHandler(w, r)
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Server is running on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
