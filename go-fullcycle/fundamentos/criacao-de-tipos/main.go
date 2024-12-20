package main

type ID int

type TIPO string

var (
	f ID = 100
	a TIPO = "Teste"
)

func main() {
	println(f)
	println(a)
}