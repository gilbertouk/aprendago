package main

import "fmt"

func main() {
	fmt.Println("Exemplo de loop for com controle de fluxo")

	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora: ", horas)

		for minutos := 0; minutos < 60; minutos++ {
			fmt.Print(" ", minutos)
		}
		fmt.Println("")
	}
	fmt.Println("Loop for concluído.")

	fmt.Println("Exemplo de loop for com break")
	forWithBreak()
	fmt.Println("Exemplo de loop for com break concluído.")

	fmt.Println("Exemplo de loop for com continue")
	forWithContinue()
	fmt.Println("Exemplo de loop for com continue concluído.")

	fmt.Println("Exemplo de loop for com saída formatada")
	forWithFormattedOutput()
	fmt.Println("Exemplo de loop for com saída formatada concluído.")
}

func forWithBreak() {
	x := 0

	for {
		if x < 1000 {
			fmt.Println("Valor de x: ", x)
			x++
		} else {
			fmt.Println("Valor de x é maior ou igual a 1000, saindo do loop.")
			break
		}

	}
}

func forWithContinue() {
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			// Se o número for ímpar, pula para a próxima iteração
			continue
		}
		fmt.Println("Número par:", i)
	}
}

func forWithFormattedOutput() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d - %v\n", i, string(rune(i)))
	}
}
