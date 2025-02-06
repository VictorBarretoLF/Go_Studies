package main

import (
	"bufio"
	"fmt"
	"os"
)

const fileDir = "arquivo.txt"

func main() {
	f, err := os.Create(fileDir)
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Hello World\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Tamanho: %d bytes\n", tamanho)
	f.Close()

	arquivo, err := os.Open(fileDir)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove(fileDir)
	if err != nil {
		panic(err)
	}
}
