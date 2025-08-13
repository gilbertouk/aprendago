package main

import "fmt"

func main() {
	executarComCallback(func() {
		fmt.Println("Executando callback!")
	})

	// para executar uma callback deve-se passar uma função como argumento
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	total := somenteNumerosImpares(somaVariosNumeros, numeros...)
	fmt.Println("Total dos numeros impares:", total)
}

func executarComCallback(callback func()) {
	callback()
}

func somaVariosNumeros(x ...int) int {
	soma := 0
	for _, numero := range x {
		soma += numero
	}
	return soma
}

func somenteNumerosImpares(f func(x ...int) int, numeros ...int) int {
	var slice []int
	for _, numero := range numeros {
		if numero%2 != 0 {
			slice = append(slice, numero)
		}
	}
	// Chama a função de callback com os números ímpares
	total := f(slice...)
	return total
}
