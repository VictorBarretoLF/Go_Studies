package main

import "fmt"

var globalVariable int = 100

var variable1 int = 200

func main() {
	var localVariable int = 200

	fmt.Printf("The value of globalVariable is %d\n", globalVariable)

	fmt.Printf("The value of localVariable is %d\n\n", localVariable)

	display()

	var variable1 int = 500
	fmt.Printf("The value of variable1 is %d\n", variable1)
}

func display() {
	fmt.Printf("The value of globalVariable is %d\n\n", globalVariable)
}
