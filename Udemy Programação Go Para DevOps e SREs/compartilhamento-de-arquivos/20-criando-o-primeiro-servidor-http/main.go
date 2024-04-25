package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	var dir string
	if runtime.GOOS == "windows" {
		dir = "C:/Users/BeHOH/Desktop"
	} else {
		dir = "/home/username/Desktop"
	}

	fs := http.FileServer(http.Dir(dir))
	fmt.Println("Subindo servidor na porta 9000!")
	log.Fatal(http.ListenAndServe(":9000", fs))
}
