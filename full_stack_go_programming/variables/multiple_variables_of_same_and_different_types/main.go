package main

import "fmt"

func main() {
	var1, var2, var3 := 100, 200, 300

	fmt.Printf("The Value of var1 is %d\n", var1)
	fmt.Printf("The type of var1 is %T\n", var1)

	fmt.Printf("The Value of var2 is %d\n", var2)
	fmt.Printf("The type of var2 is %T\n", var2)

	fmt.Printf("The Value of var3 is %d\n", var3)
	fmt.Printf("The type of var3 is %T\n", var3)

	fmt.Printf("\n\n")

	var4, var5, var6 := 100, "Go programming", 3.14

	fmt.Printf("The Value of var4 is %d\n", var4)
	fmt.Printf("The type of var4 is %T\n", var4)

	fmt.Printf("The Value of var5 is %s\n", var5)
	fmt.Printf("The type of var5 is %T\n", var5)

	fmt.Printf("The Value of var6 is %f\n", var6)
	fmt.Printf("The type of var6 is %T\n", var6)

}
