package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"João":  25,
		"Maria": 30,
		"José":  22,
	}

	fmt.Println("Mapa de amigos:", amigos)
	fmt.Println("Tamanho do mapa:", len(amigos))

	// adiciona gil ao mapa
	amigos["Gil"] = 28

	gil, ok := amigos["Gil"] // Verifica se a chave "Gil" existe no mapa
	if ok {
		fmt.Println("Idade de Gil:", gil)
	} else {
		fmt.Println("Gil não encontrado no mapa.")
	}

	fmt.Println()

	// iteração sobre o mapa com range
	for nome, idade := range amigos {
		fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
	}

	fmt.Println()

	totalIdade := 0

	for _, idade := range amigos {
		totalIdade += idade
	}

	fmt.Println("Idade total:", totalIdade)
	fmt.Println("Idade média:", totalIdade/len(amigos))

	fmt.Println()

	// removendo um amigo do mapa
	delete(amigos, "José") // Remove a chave "José" do mapa
	fmt.Println("Mapa após remoção de José:", amigos)
}
