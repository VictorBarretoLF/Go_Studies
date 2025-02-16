package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome        string
	Idade       int
	Ativo       bool
	Localizacao Endereco
}

func main() {
	victor := Cliente{
		Nome:  "Victor",
		Idade: 29,
		Ativo: true,
	}

	victor.Ativo = false
	victor.Localizacao.Cidade = "cidade"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", victor.Nome, victor.Idade, victor.Ativo)
}
