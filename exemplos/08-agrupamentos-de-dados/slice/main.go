package main

import (
	"fmt"
)

func main() {
	slice := []int{}

	for i := 1; i <= 32; i++ {
		slice = append(slice, i)
	}

	fmt.Println("Slice:", slice)
	fmt.Println("Tamanho do slice:", len(slice))
	fmt.Println("Capacidade do slice:", cap(slice))

	fmt.Println()

	names := []string{"João", "Maria", "José", "Ana", "Pedro"}
	// names[5] = "Lucas" // Isso causará um erro de índice fora do intervalo

	for index, name := range names {
		fmt.Printf("Índice: %d, Nome: %s\n", index, name)
	}

	fmt.Println()
	// Exemplo de iteração com índice
	for index := 0; index < len(names); index++ {
		fmt.Printf("Índice: %d, Nome: %s\n", index, names[index])
	}

	fmt.Println()
	// slice de um slice
	subSlice := names[1:] // Pega os elementos do índice 1 ao final do slice
	// subSlice := names[1:len(names)] // Outra forma de fazer o mesmo
	// subSlice := names[1:5] // Pega os elementos do índice 1 ao 4
	fmt.Println("Sub-slice:", subSlice)

	fmt.Println()

	// removendo elementos de um slice
	// Exemplo: remover o elemento no índice 2 (José)
	names = append(names[:2], names[3:]...) // Remove o elemento no índice 2
	fmt.Println("Slice após remoção do elemento no índice 2 'José':", names)

	fmt.Println()

	// Adicionando elementos ao slice
	outrosNomes := []string{"Lucas", "Fernanda", "Carlos"}

	names = append(names, outrosNomes...) // Adiciona os novos nomes ao final do slice
	// names = append(names, "Lucas", "Fernanda", "Carlos") // Outra forma de adicionar
	fmt.Println("Slice após adicionar novos nomes:", names)

	fmt.Println()

	// Exemplo de slice com make
	sliceMake := make([]int, 5, 10) // Cria um slice com 5 elementos e capacidade para 10
	sliceMake[0] = 1
	sliceMake[1] = 2
	sliceMake[2] = 3
	sliceMake[3] = 4
	sliceMake[4] = 5
	fmt.Println("Slice criado com make:", sliceMake)
	fmt.Println("Tamanho do slice criado com make:", len(sliceMake))
	fmt.Println("Capacidade do slice criado com make:", cap(sliceMake))

	// Ultrapassando a capacidade do slice
	// Adiciona mais elementos, ultrapassando a capacidade inicial
	// ele aumentará automaticamente a capacidade para o dobro definido inicialmente
	sliceMake = append(sliceMake, 6, 7, 8, 9, 10, 11)
	fmt.Println("Slice após ultrapassar a capacidade:", sliceMake)
	fmt.Println("Novo tamanho do slice:", len(sliceMake))
	fmt.Println("Nova capacidade do slice:", cap(sliceMake))

	fmt.Println()

	// Exemplo de slice multi-dimensional
	sliceMulti := [][]string{
		{"João", "Maria"},
		{"José", "Ana"},
		{"Pedro", "Lucas"},
	}

	fmt.Println("Slice multi-dimensional:", sliceMulti)
}
