package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, please send an email to <a href=\"mailto:victorbarretolins@gmail.com\">victorbarretolins@gmail.com</a>.</p>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("Server is running on port 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
