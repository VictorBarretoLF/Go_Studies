package main

import (
	"fmt"
	"gohello/packages/funcionarios"
)

func main() {
	pessoa := &funcionarios.Pessoa{
		Nome:    "Victor",
		Idade:   39,
		Salario: 100,
	}

	fmt.Println(pessoa)
}
