package main

import "fmt"

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
}

func main()  {
	victor := Cliente{
		Nome: "Victor",
		Idade: 29,
		Ativo: true,
	}

	victor.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", victor.Nome, victor.Idade, victor.Ativo)
}