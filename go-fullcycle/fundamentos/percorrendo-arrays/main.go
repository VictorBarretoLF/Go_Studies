package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(len(meuArray))
	
	for i, v := range meuArray {
		fmt.Printf("Index=%d, Value=%d\n",i, v)
	}
}