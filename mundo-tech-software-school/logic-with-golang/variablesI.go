package main

import "fmt"

func main() {
	// Se não definir o valor da variável, o Go vai iniciá-la como zero - no caso dos inteiros.
	// O tipo foi declarado
	var age int
	var newAge int = 23
	var name string
	name = "Abraão Santos"
	fmt.Println(age)
	fmt.Println(newAge)
	fmt.Println(name)

	// Também tem como fazer inferência de tipos
	// Ou seja, o compilador define o tipo conforme estabelecemos
	newName := "Abraão"
	fmt.Println(newName)

	var ageOne, ageTwo int = 20, 30
	fmt.Println("Idades com variáveis combinadas: ")
	fmt.Println(ageOne, ageTwo)
}