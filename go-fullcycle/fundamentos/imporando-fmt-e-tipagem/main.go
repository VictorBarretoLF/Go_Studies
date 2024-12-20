package main

import "fmt"

const a = "Hello World"

type ID int

var (
	b ID = 100
	c string = "Teste"
	d bool = true
	e float64 = 1.5
)

func main() {
	fmt.Printf("O tipo de b é : %T\n", b)
	fmt.Printf("O tipo de c é : %T\n", c)
	fmt.Printf("O tipo de d é : %T\n", d)
	fmt.Printf("O tipo de e é: %T\n", e)
	fmt.Printf("O tipo de e é: %v\n", e)
}