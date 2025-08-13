package main

import "fmt"

func main() {
	// defer: executa essa funcao ao final, o primeiro defer e o ultimo a ser executado
	defer fmt.Println("Fim do programa")

	soma := func(a int, b int) int {
		return a + b
	}

	subtracao := func(a int, b int) int {
		return a - b
	}

	fmt.Println("Soma:", soma(5, 3))
	fmt.Println("Subtração:", subtracao(5, 3))

	numeros := []int{1, 2, 3, 4, 5}

	total := somaVariosNumeros(numeros...)
	fmt.Println("Soma de Vários Números:", total)

	pessoa := Pessoa{
		Nome:  "Gilberto Silva",
		Idade: 30,
	}
	pessoa.imprimirDados()
}

func somaVariosNumeros(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

// Cria um tipo pessoa
type Pessoa struct {
	Nome  string
	Idade int
}

// Cria uma função que adiciona um método para o tipo Pessoa
func (p Pessoa) imprimirDados() {
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Idade:", p.Idade)

	fmt.Println(p.Nome, "diz bom dia")
}
