package main

import "fmt"

func main() {
	/*fmt.Println()
	fmt.Println("Olá GO")
	fmt.Println()
	fmt.Println(5+5)
	fmt.Println("Variáveis do tipo String")

	*/var nome string = "Carlos"
	var sobrenome string = "Augusto Petri da Silva"
	var idade int = 20

	fmt.Println("O Aluno " + nome, sobrenome, "tem", idade, "anos")
    //int8
	var numero1 int8 = 10
	var numero2 int8 = 5

    soma := int(numero1) + int(numero2)
	multiplicacao := int(numero1) * int(numero2)
	divisao := int(numero1) / int(numero2)

	fmt.Println(soma)
	fmt.Println(multiplicacao)
	fmt.Println(divisao)

	numero1++
	fmt.Println(numero1)
	
	numero1--
	// numero1 = numero1 - 1
	numero1--
	numero1--
	fmt.Println(numero1)
}

