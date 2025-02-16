package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

// Method with a pointer receiver to modify the original struct
func (c *Cliente) DesativarComPonteiro() {
	c.Ativo = false
}

// Method with a value receiver (works on a copy of the struct)
func (c Cliente) DesativarSemPonteiro() {
	c.Ativo = false
}

func main() {
	victor := Cliente{
		Nome:  "Victor",
		Idade: 29,
		Ativo: true,
	}

	// Using the pointer receiver method to modify the original struct
	victor.DesativarComPonteiro()
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", victor.Nome, victor.Idade, victor.Ativo) // Output: Nome: Victor, Idade: 29, Ativo: false

	// Reset Ativo to true for demonstration
	victor.Ativo = true

	// Using the value receiver method (does not modify the original struct)
	victor.DesativarSemPonteiro()
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", victor.Nome, victor.Idade, victor.Ativo) // Output: Nome: Victor, Idade: 29, Ativo: true
}