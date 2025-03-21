package main

import "fmt"

func main() {
	// Declaração de variáveis <-> <var> <nome_variavel> <tipo>
	var helloWorldOne string
	var helloWorldTwo string

	fmt.Println("Digite 'OláMundo!'")
	fmt.Scan(&helloWorldOne) // Lê até o primeiro espaço.

	if helloWorldOne == "OláMundo!" { // Condicional no Golang
		fmt.Println("Correto.")
	} else {
		fmt.Println("Incorreto.")
	}

	fmt.Println("Digite a frase 'Olá, mundo!'")
	fmt.Println()
	fmt.Scanln(&helloWorldTwo)

	if helloWorldTwo == "Olá, mundo!" {
		fmt.Println("Correto.")
	} else {
		fmt.Println("Incorreto.")
	}

}
