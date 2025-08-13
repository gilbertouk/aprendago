package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  int
	Cidade string
}

type Desenvolvedor struct {
	Pessoa
	Linguagem string
	Salario   float64
	AnosExp   int
}

type Gerente struct {
	Pessoa
	Departamento string
	Salario      float64
}

func (c Desenvolvedor) ImprimirDados() {
	fmt.Printf("Desenvolvedor: %s, Idade: %d, Cidade: %s, Linguagem: %s, Salário: %.2f, Anos de Experiência: %d\n",
		c.Nome, c.Idade, c.Cidade, c.Linguagem, c.Salario, c.AnosExp)
}

func (c Gerente) ImprimirDados() {
	fmt.Printf("Gerente: %s, Idade: %d, Cidade: %s, Departamento: %s, Salário: %.2f\n",
		c.Nome, c.Idade, c.Cidade, c.Departamento, c.Salario)
}

type Imprimivel interface {
	ImprimirDados()
}

func ImprimirDadosDePessoa(i Imprimivel) {
	// Chama o método ImprimirDados da interface
	i.ImprimirDados()

	// Podemos ter um switch aqui para tratar diferentes tipos
	switch v := i.(type) {
	case Desenvolvedor:
		fmt.Println("Tipo: Desenvolvedor")
		// Acessar campos específicos de Desenvolvedor
		fmt.Printf("Linguagem: %s\n", v.Linguagem)
	case Gerente:
		fmt.Println("Tipo: Gerente")
		// Acessar campos específicos de Gerente
		fmt.Printf("Departamento: %s\n", v.Departamento)
	}
}

func main() {
	dev := Desenvolvedor{
		Pessoa:    Pessoa{"Alice", 30, "São Paulo"},
		Linguagem: "Go",
		Salario:   80000,
		AnosExp:   5,
	}
	ImprimirDadosDePessoa(dev)

	ger := Gerente{
		Pessoa:       Pessoa{"Bob", 40, "Rio de Janeiro"},
		Departamento: "TI",
		Salario:      120000,
	}
	ImprimirDadosDePessoa(ger)
}
