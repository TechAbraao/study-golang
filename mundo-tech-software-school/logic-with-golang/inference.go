package main

import "fmt"

func main() {
	var i int
	var j = i
	// Se o valor não foi iniciado, vai iniciar como zero
	fmt.Printf("\nTipo: %T e Valor: %v\n", i, j)

	// O próprio Golang infere o tipo de dado
	var d = 36
	var e = 36.00
	fmt.Printf("\nTipo: %T e Valor: %v\n", d, d)
	fmt.Printf("\nTipo: %T e Valor: %v\n", e, d)
}