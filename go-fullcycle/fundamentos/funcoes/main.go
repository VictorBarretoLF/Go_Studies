package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sumIsGreaterThan50(1, 2))

	valor, err := sumIsGreaterThan50(30, 50)
	if err != nil {
		fmt.Println("Erro aqui", err)
	}
	fmt.Println(valor)
}

func sum(a int, b int) int {
	return a + b
}

func sumIsGreaterThan50(a int, b int) (int, error) {
	if a + b >= 50 {
		return a + b, errors.New("O valor é maior que 50")
	}
	return a + b, nil
}
