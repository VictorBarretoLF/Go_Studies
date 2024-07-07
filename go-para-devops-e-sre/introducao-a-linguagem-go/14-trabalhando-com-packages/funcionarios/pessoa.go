package funcionarios

/*
 * The error you're encountering, "unknown field 'nome' in struct literal of type funcionarios.Pessoa",
 * occurs because the fields in your Pessoa struct are unexported (private) due to them starting with a
 * lowercase letter. In Go, identifiers that start with a lowercase letter are only accessible within
 * the same package. To access these fields from another package, you need to export them by starting
 * their names with an uppercase letter.
 */
type Pessoa struct {
	Nome    string
	Idade   int
	Salario int
}

func (p *Pessoa) AddSalary(bonus int) {
	p.Salario += bonus
}
