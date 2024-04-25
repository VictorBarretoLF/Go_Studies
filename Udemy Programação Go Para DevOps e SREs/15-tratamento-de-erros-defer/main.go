package main

import (
	"fmt"
	"os"
	"strconv"
)

// # Rode o comando `go run 10`
// # Rode o comando `go run hello`
func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	if len(os.Args) != 2 {
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("erro, não é um número válido")
		os.Exit(1)
	}
	fmt.Println("Numero convertido: ", n)
}
