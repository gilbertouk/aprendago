package main

import "fmt"

func main() {
	soma := func(a int, b int) int {
		return a + b
	}

	subtracao := func(a int, b int) int {
		return a - b
	}

	fmt.Println("Soma:", soma(5, 3))
	fmt.Println("Subtração:", subtracao(5, 3))
}
