package main

import "fmt"

type Empregado struct {
	Nome string
	id   int
}

func main() {
	emp := Empregado{"Victor", 743}
	pts := &emp

	fmt.Println(pts) // &{Victor 743}
	pts.Nome = "Lucas"
	fmt.Println(emp) // {Lucas 743}
	fmt.Println(pts) // &{Lucas 743}

	noPointerFunction(emp)
	fmt.Println("Printando após chamada da função sem ponteiro", emp) // {Lucas 743}

	pointerFunction(&emp)
	fmt.Println("Printando após chamada da função *com* ponteiro", emp) // {Marcus 743}

}

func pointerFunction(e *Empregado) {
	fmt.Println("Chamada da funçã", e)
	e.Nome = "Marcus"
}

func noPointerFunction(e Empregado) {
	fmt.Println("Chamada da funçã", e)
	e.Nome = "Marcus"
}
