package main

import "fmt"

func main() {
	var variable1, variable2, variable3 = 2, "John", 1.75

	fmt.Printf("variable1 = %d\n", variable1)
	fmt.Printf("The type of variable1 is %T\n", variable1)

	fmt.Printf("variable2 = %s\n", variable2)
	fmt.Printf("The type of variable2 is %T\n", variable2)

	fmt.Printf("variable3 = %f\n", variable3)
	fmt.Printf("The type of variable3 is %T\n", variable3)
}
