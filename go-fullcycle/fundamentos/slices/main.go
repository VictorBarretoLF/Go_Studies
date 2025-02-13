package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)  
	// len=9 cap=9 [10 20 30 40 50 60 70 80 90] (slice original)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])  
	// len=0 cap=9 [] (slice vazio, mas mantém a capacidade original)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])   
	// len=4 cap=9 [10 20 30 40] (4 primeiros elementos, capacidade mantida)

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])   
	// len=7 cap=7 [30 40 50 60 70 80 90] (subslice a partir do índice 2, capacidade reduzida)
	
	s = append(s, 110)  // Adiciona 110 (len=10, cap=18 - dobra a capacidade)
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])   
	// len=8 cap=16 [30 40 50 60 70 80 90 110] (nova capacidade: 18 - 2 = 16)
	
	s = append(s, 120)  // Adiciona 120 (len=11, cap=18 - sem realocação)
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])   
	// len=9 cap=16 [30 40 50 60 70 80 90 110 120] (mesma capacidade do array subjacente)
}