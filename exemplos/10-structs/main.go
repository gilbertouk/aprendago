package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	pessoa1 := Pessoa{"Alice", "Silva", false}
	pessoa2 := Pessoa{"Bob", "Santos", true}
	pessoa3 := Pessoa{
		nome:      "Carol",
		sobrenome: "Oliveira",
		fumante:   false,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa3)
}
