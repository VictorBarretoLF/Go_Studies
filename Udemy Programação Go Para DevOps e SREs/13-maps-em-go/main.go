package main

import (
	"fmt"
)

type Pessoa struct {
	name    string
	salario int
}

func main() {
	salFuncs := make(map[string]int)
	salFuncs["Victor"] = 10
	salFuncs["Genivaldo"] = 30

	sal, exists := salFuncs["Luiz"]
	salarioVictor, salarioVictorExist := salFuncs["Victor"]

	fmt.Println(sal, exists) // 0, false
	fmt.Println(salarioVictor, salarioVictorExist) // 10, true

	totalSal := len(salFuncs)
	fmt.Println(totalSal)

}
