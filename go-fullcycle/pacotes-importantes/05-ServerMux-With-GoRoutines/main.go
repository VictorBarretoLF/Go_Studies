package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World"))
	// })

	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})

	go func() {
		fmt.Println("Starting server on :3000")
		err := http.ListenAndServe(":3000", mux)
		if err != nil {
			fmt.Printf("Error starting server on :3000: %s\n", err)
		}
	}()

	// A parte abaixo não é hitada
	fmt.Println("Hitando aqui")

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Accessing the mux 2")
		w.Write([]byte("Hello World 2"))
	})

	go func() {
		fmt.Println("Starting server on :9000")
		if err := http.ListenAndServe(":9000", mux2); err != nil {
			fmt.Printf("Error starting server on :9000: %s\n", err)
		}
	}()

	select {}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Blog"))
}
