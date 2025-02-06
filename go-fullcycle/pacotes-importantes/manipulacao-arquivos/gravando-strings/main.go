package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Hello World\n")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Tamanho: %d bytes\n", tamanho)

	f.Close()
}