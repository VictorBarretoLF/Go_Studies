package main

import "fmt"

func pointerFunction(a *int) {
	*a = 400
}

func main() {
	var x = 100
	fmt.Println("o valor da variável X antes da função é:", x)

	var pa *int = &x

	pointerFunction(pa)
	fmt.Println("o valor da variável X depois da função é:", x)
}