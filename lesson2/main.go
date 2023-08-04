package main

import "fmt"

// you cannot use:
// var someName := "hello"
// outside of a function

func main() {

	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string // it defaults to an empty string

	fmt.Println(nameOne, nameTwo, nameThree) // mario luigi

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree) // peach luigi bowser

	nameFour := "yoshi"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	// docs: https://pkg.go.dev/builtin#int
	var numOne int8 = 25 // you cannot use 255 because it int8 does not suport it in an specif range
	var numTwo int8 = -128
	var numThree uint = 25

	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 12333333333333333321123123.7
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
