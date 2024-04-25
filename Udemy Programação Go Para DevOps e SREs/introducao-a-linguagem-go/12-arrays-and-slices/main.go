package main

import (
	"fmt"
)

type Pessoa struct {
	name    string
	idade   int
	salario int
}

func main() {
	salarios := []int{}

	for i := 0; i < 10; i++ {
		salarios = append(salarios, 100+i)
	}

	for index, salario := range salarios {
		fmt.Println(index, salario)
	}

	fmt.Println(salarios)
}
