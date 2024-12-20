package main

// nao posso mudar o valor de uma constante, é equivalente ao final do java
const f = 100

// posso atribuir valores a variaveis sem tipar, o go vai inferir o tipo
var (
	a = "Teste"
	b int
	c string
	d bool
	e float64
)

func main() {
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)

	// estou atribuindo e inferindo o tipo da variavel

	g := "hello"
	println(g)
}
