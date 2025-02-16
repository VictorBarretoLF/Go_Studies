package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func swapNormal(a int, b int) {
	temp := a
	a = b
	b = temp
}

func main()  {
	x, y := 10, 20
    fmt.Println("Before swap:", x, y)
	swap(&x, &y)
    fmt.Println("After swap:", x, y)

	fmt.Println("Before swap normal:", x, y)
	swapNormal(x, y)
	fmt.Println("After swap normal:", x, y)
}
