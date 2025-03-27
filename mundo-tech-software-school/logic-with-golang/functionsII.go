package main

import "fmt"

// Retornando dois valores no Golang
func calcular(a int) (int, int) {
	var quadrado int = a * a
	var cubo int = a * a * a

	return quadrado, cubo
}

// Também há como fazer o retorno ter seus nomes, isto é, valores de retorno nomeados
// Com o retorno nomeado, as variáveis a serem retornadas estão na ASSINATURA da função
// Ou seja, evitamos a instrução 'return'. Veja exemplos:

// Aqui especifica o que será retornado.
func nomeSobrenome(nome, sobrenome string) (nomeCompleto string) {
	nomeCompleto = nome + " " + sobrenome // A função retornará o 'nomeCompleto'
	return
}

func main() {
	fmt.Println(calcular(2))
	fmt.Println(nomeSobrenome("José", "Teixeira"))
}
