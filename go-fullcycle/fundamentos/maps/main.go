package main

import "fmt"

func main() {
    salarios := map[string]int{"Victor": 100, "John": 200, "Lucas": 300}
    
    // Prints 100 (value of "Victor" before deletion)
    fmt.Println(salarios["Victor"])
    
    delete(salarios, "Victor")
    salarios["Pedro"] = 400
    
    // Prints map[John:200 Lucas:300 Pedro:400] (order may vary)
    fmt.Println(salarios)

    salariosV2 := make(map[string]int)
    salariosV2["Thiago"] = 100
    
    // Prints map[Thiago:100] (only entry in new map)
    fmt.Println(salariosV2)

    // Prints separator line
    fmt.Println("-------------")

    // Prints salaries with names in random order. Example output:
    // O salario de John é 200
    // O salario de Lucas é 300
    // O salario de Pedro é 400
    for nome, salario := range salarios {
        fmt.Printf("O salario de %s é %d\n", nome, salario)
    }

    // Prints separator line
    fmt.Println("-------------")

    // Prints salaries without names in DIFFERENT random order. Example output:
    // O salario é de 200
    // O salario é de 300
    // O salario é de 400
    for _, salario := range salarios {
        fmt.Printf("O salario é de %d\n", salario)
    }
}