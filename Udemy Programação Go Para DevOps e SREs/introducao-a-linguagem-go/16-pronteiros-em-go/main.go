package main

import "fmt"

func main() {
	var inteiro = 45
	var ponteiro *int = &inteiro

	fmt.Println("Valor do inteiro:", inteiro)
	fmt.Println("Endereço da variavel inteiro:", &inteiro)
	fmt.Println("Valor da variavel ponteiro:", ponteiro)
	fmt.Println("Endereço da variavel ponteiro:", &ponteiro)
	fmt.Println("Endereço da variavel inteiro pelo ponteiro:", *ponteiro)
}
