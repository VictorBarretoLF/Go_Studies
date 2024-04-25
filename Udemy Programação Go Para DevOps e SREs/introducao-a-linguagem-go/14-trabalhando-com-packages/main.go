package main

import (
	"fmt"
	"gotraining/14-trabalhando-com-packages/funcionarios"
)

func main() {
	pessoa := &funcionarios.Pessoa{
		Nome:    "Victor",
		Idade:   39,
		Salario: 100,
	}

	fmt.Println(pessoa)
}
