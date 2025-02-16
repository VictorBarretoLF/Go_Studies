package main

import "fmt"

func main() {
    // Calling the variadic function `sum` with different numbers of arguments
    fmt.Println(sum(1, 2, 3))         // Calls `sum` with arguments 1, 2, 3 and prints the result (Output: 6)
    fmt.Println(sum(10, 20, 30, 40))  // Calls `sum` with arguments 10, 20, 30, 40 and prints the result (Output: 100)
    fmt.Println(sum())                // Calls `sum` with no arguments and prints the result (Output: 0)

    // Declaring and immediately invoking an anonymous function (closure)
    total := func() int {
        // Calls the `sum` function with arguments 1, 3, 4, 5, 6, 7, 8, 9, 10
        return sum(1, 3, 4, 5, 6, 7, 8, 9, 10)
    }()  // The `()` at the end immediately invokes the anonymous function
    fmt.Println(total)  // Prints the result returned by the anonymous function

	// Create a closure
	increment := counter()

	// Call the closure multiple times
	fmt.Println(increment()) // Output: 1
	fmt.Println(increment()) // Output: 2
	fmt.Println(increment()) // Output: 3

	// Create a new closure (it has its own state)
	newIncrement := counter()
	fmt.Println(newIncrement()) // Output: 1 (independent state)
}

// Variadic function to calculate the sum of integers
func sum(numbers ...int) int {
    total := 0  // Initializes a variable to store the sum
    for _, num := range numbers {  // Iterates over each number in the `numbers` slice
        total += num  // Adds the current number to the `total`
    }
    return total  // Returns the computed sum
}

// A closure is a function that references variables from outside its body
func counter() func() int {
    count := 0 // This variable is "captured" by the closure
    return func() int {
        count++ // The inner function modifies the captured variable
        return count
    }
}