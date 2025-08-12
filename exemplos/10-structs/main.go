package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	fumante   bool
}

type Profissional struct {
	Pessoa
	cargo   string
	salario float64
}

func main() {
	pessoa1 := Pessoa{"Alice", "Silva", false}
	pessoa2 := Pessoa{"Bob", "Santos", true}
	pessoa3 := Pessoa{
		nome:      "Carol",
		sobrenome: "Oliveira",
		fumante:   false,
	}
	profissional1 := Profissional{
		Pessoa:  Pessoa{"Daniel", "Pereira", false},
		cargo:   "Desenvolvedor",
		salario: 5000.00,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa3)
	fmt.Println(profissional1)

	// struct anônimos
	x := struct {
		titulo  string
		salario float64
	}{
		titulo:  "Desenvolvedor",
		salario: 5000.00,
	}

	fmt.Println(x)
}
