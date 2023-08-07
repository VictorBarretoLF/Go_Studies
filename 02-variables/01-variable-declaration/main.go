package main

import "fmt"

func main() {
	// docs: https://pkg.go.dev/builtin#int
	// Declaring Variables
	var myStr string = "Hello"
	var myInt int = 100
	var myFloat float64 = 45.12
	fmt.Println(myStr, myInt, myFloat) // Hello 100 45.12

	//================================

	// Multiple Declarations
	var (
		employeeId          int    = 5
		firstName, lastName string = "Satoshi", "Nakamoto"
	)
	fmt.Println(employeeId, firstName, lastName) // 5 Satoshi Nakamoto

	//================================

	// Short variable declaration syntax
	name := "Rajeev Singh"
	age, salary, isProgrammer := 35, 50000.0, true

	fmt.Println(name, age, salary, isProgrammer) // Rajeev Singh 35 50000 true
}
