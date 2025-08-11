package main

import "fmt"

func main() {
	fmt.Println("Exercícios de Agrupamentos de Dados")
	fmt.Println("================================")
	fmt.Println("Exercício 1:")
	exercicio1()
	fmt.Println("\nExercício 2:")
	exercicio2()
	fmt.Println("\nExercício 3:")
	exercicio3()
	fmt.Println("\nExercício 4:")
	exercicio4()
	fmt.Println("\nExercício 5:")
	exercicio5()
	fmt.Println("\nExercício 6:")
	exercicio6()
	fmt.Println("\nExercício 7:")
	exercicio7()
	fmt.Println("\nExercício 8:")
	exercicio8()
	fmt.Println("\nExercício 9:")
	exercicio9()
	fmt.Println("\nExercício 10:")
	exercicio10()
}

func exercicio1() {
	//	Usando uma literal composta:
	//
	// Crie um array que suporte 5 valores to tipo int
	// Atribua valores aos seus índices
	// Utilize range e demonstre os valores do array.
	// Utilizando format printing, demonstre o tipo do array.

	array := [5]int{10, 20, 30, 40, 50}

	for i, v := range array {
		fmt.Println("Índice:", i, "Valor:", v)
	}

	fmt.Printf("Tipo do array: %T\n", array)
}

func exercicio2() {
	// 	Usando uma literal composta:
	// Crie uma slice de tipo int
	// Atribua 10 valores a ela
	// Utilize range para demonstrar todos estes valores.
	// E utilize format printing para demonstrar seu tipo.

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range slice {
		fmt.Println("Índice:", i, "Valor:", v)
	}
	fmt.Printf("Tipo da slice: %T\n", slice)
}

func exercicio3() {
	//	Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
	// Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
	// Do quinto ao último item do slice (incluindo o último item!)
	// Do segundo ao sétimo item do slice (incluindo o sétimo item!)
	// Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
	// Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Do primeiro ao terceiro item:", slice[0:3])
	fmt.Println("Do quinto ao último item:", slice[4:])
	fmt.Println("Do segundo ao sétimo item:", slice[1:7])
	fmt.Println("Do terceiro ao penúltimo item:", slice[2:9])
	// Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
	fmt.Println("Do terceiro ao penúltimo item (usando len()):", slice[2:len(slice)-1])
}

func exercicio4() {
	//	Começando com a seguinte slice:
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// Anexe a ela o valor 52;
	// Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
	// Demonstre a slice;
	// Anexe a ela a seguinte slice:
	// y := []int{56, 57, 58, 59, 60}
	// Demonstre a slice x.

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	fmt.Println("Slice após anexar 52, 53, 54 e 55:", x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...) // Anexa a slice y a x
	fmt.Println("Slice x após anexar y:", x)
}

func exercicio5() {
	//	Comece com essa slice:	//
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// Utilizando slicing e append, crie uma slice y que contenha os valores:
	// [42, 43, 44, 48, 49, 50, 51]

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println("Slice x:", x)
}

func exercicio6() {
	//	Crie uma slice usando make que possa conter todos os estados do Brasil.	//
	// Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
	// Demonstre o len e cap da slice.
	// Demonstre todos os valores da slice sem utilizar range.

	estados := make([]string, 0)
	estados = append(estados, "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins")

	fmt.Println("Quantidade de estados:", len(estados))
	fmt.Println("Capacidade da slice:", cap(estados))
	fmt.Println("Estados do Brasil:")
	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}
}

func exercicio7() {
	// Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
	// "Nome", "Sobrenome", "Hobby favorito"
	// Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

	pessoas := [][]string{
		{"Alice", "Silva", "Leitura"},
		{"Bob", "Santos", "Futebol"},
		{"Carol", "Oliveira", "Cozinhar"},
	}

	for i, pessoa := range pessoas {
		fmt.Printf("Pessoa %d:\n", i+1)
		for j, dado := range pessoa {
			fmt.Printf("  Dado %d: %s\n", j+1, dado)
		}
	}
}

func exercicio8() {
	//	Crie um map com key tipo string e value tipo []string.
	// Key deve conter nomes no formato sobrenome_nome
	// Value deve conter os hobbies favoritos da pessoa
	// Demonstre todos esses valores e seus índices.

	hobbies := make(map[string][]string)

	hobbies["Silva_Alice"] = []string{"Leitura"}
	hobbies["Santos_Bob"] = []string{"Futebol"}
	hobbies["Oliveira_Carol"] = []string{"Cozinhar"}

	for nome, hobby := range hobbies {
		fmt.Printf("Nome: %s\n", nome)
		for i, h := range hobby {
			fmt.Printf("  Hobby %d: %s\n", i+1, h)
		}
	}
}

func exercicio9() {
	// Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.

	hobbies := make(map[string][]string)

	hobbies["Silva_Alice"] = []string{"Leitura"}
	hobbies["Santos_Bob"] = []string{"Futebol"}
	hobbies["Oliveira_Carol"] = []string{"Cozinhar"}

	hobbies["Souza_Daniel"] = []string{"Caminhada"}

	for nome, hobby := range hobbies {
		fmt.Printf("Nome: %s\n", nome)
		for i, h := range hobby {
			fmt.Printf("  Hobby %d: %s\n", i+1, h)
		}
	}
}

func exercicio10() {
	// Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.

	hobbies := make(map[string][]string)
	hobbies["Silva_Alice"] = []string{"Leitura"}
	hobbies["Santos_Bob"] = []string{"Futebol"}
	hobbies["Oliveira_Carol"] = []string{"Cozinhar"}
	hobbies["Souza_Daniel"] = []string{"Caminhada"}

	for nome, hobby := range hobbies {
		fmt.Printf("Nome: %s\n", nome)
		for i, h := range hobby {
			fmt.Printf("  Hobby %d: %s\n", i+1, h)
		}
	}
}
