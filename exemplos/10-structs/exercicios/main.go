package main

import "fmt"

func main() {
	fmt.Println("Exercício 1:")
	exercicio1()
	fmt.Println()
	fmt.Println("Exercício 2:")
	exercicio2()
	fmt.Println()
	fmt.Println("Exercício 3:")
	exercicio3()
	fmt.Println()
	fmt.Println("Exercício 4:")
	exercicio4()
}

// Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
// Nome
// Sobrenome
// Sabores favoritos de sorvete
// Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func exercicio1() {
	pessoa1 := pessoa{
		nome:      "Alice",
		sobrenome: "Silva",
		sabores:   []string{"chocolate", "baunilha", "morango"},
	}
	pessoa2 := pessoa{
		nome:      "Bob",
		sobrenome: "Santos",
		sabores:   []string{"flocos", "limão", "coco"},
	}

	fmt.Println("Sabores favoritos de", pessoa1.nome, pessoa1.sobrenome)
	for _, sabor := range pessoa1.sabores {
		fmt.Println("-", sabor)
	}

	fmt.Println("Sabores favoritos de", pessoa2.nome, pessoa2.sobrenome)
	for _, sabor := range pessoa2.sabores {
		fmt.Println("-", sabor)
	}
}

// Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// Demonstre os valores do map utilizando range.
// Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.

func exercicio2() {
	pessoas := make(map[string]pessoa)

	pessoas["Silva"] = pessoa{
		nome:      "Alice",
		sobrenome: "Silva",
		sabores:   []string{"chocolate", "baunilha", "morango"},
	}
	pessoas["Santos"] = pessoa{
		nome:      "Bob",
		sobrenome: "Santos",
		sabores:   []string{"flocos", "limão", "coco"},
	}

	for sobrenome, p := range pessoas {
		fmt.Println("Sabores favoritos de", p.nome, sobrenome)
		for _, sabor := range p.sabores {
			fmt.Println("-", sabor)
		}
	}
}

// Crie um novo tipo: veículo
//  - O tipo subjacente deve ser struct
//  - Deve conter os campos: portas, cor
// Crie dois novos tipos: caminhonete e sedan
//  - Os tipos subjacentes devem ser struct
//  - Ambos devem conter "veículo" como struct embutido
//  - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
//  - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
// Usando os structs veículo, caminhonete e sedan:
//  - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
//  - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
// Demonstre estes valores.
// Demonstre um único campo de cada um dos dois.

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func exercicio3() {
	caminhonete := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "azul",
		},
		tracaoNasQuatro: true,
	}

	sedan := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "preto",
		},
		modeloLuxo: true,
	}

	fmt.Println("Caminhonete:")
	fmt.Println("Portas:", caminhonete.portas)
	fmt.Println("Cor:", caminhonete.cor)
	fmt.Println("Tração nas quatro:", caminhonete.tracaoNasQuatro)

	fmt.Println("\nSedan:")
	fmt.Println("Portas:", sedan.portas)
	fmt.Println("Cor:", sedan.cor)
	fmt.Println("Modelo luxo:", sedan.modeloLuxo)
}

// Crie e use um struct anônimo.
// Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.

func exercicio4() {
	x := struct {
		nome     string
		sabores  []string
		traducao map[string]string
	}{
		nome:    "Alice",
		sabores: []string{"chocolate", "baunilha", "morango"},
		traducao: map[string]string{
			"chocolate": "chocolate",
			"baunilha":  "vanilla",
			"morango":   "strawberry",
		},
	}

	fmt.Println("Nome:", x.nome)
	fmt.Println("Sabores favoritos:")
	for _, sabor := range x.sabores {
		fmt.Println("-", sabor)
	}
	fmt.Println()
	fmt.Println("Tradução dos sabores:")
	for sabor, traducao := range x.traducao {
		fmt.Println("-", sabor, ":", traducao)
	}
}
