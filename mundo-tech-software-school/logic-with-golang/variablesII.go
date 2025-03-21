package main 

import (
	"fmt"
	"strconv" // Pra converter pro tipo String
)

func main() {
	var ( // Declarando várias variáveis em conjunto
		age int = 23
		name string = "Abraão"
	)

	fmt.Println(age)
	fmt.Println(name)
	fmt.Println("Your name is " + name + " and you has " + strconv.Itoa(age) + " years old")
	fmt.Println("Name: ", name, " and age: ", age, " years old.")


	// Pra não declarar o tipo e o Golang mesmo assim tipar de forma dinâmica
	height := 1.70
	x1, x2, x3, x4 := 10, 20, 30, 40
	numberOne, numberTwo, yourName := 20, 30, "Name is here!"
	result := (x1 + x2 + x3 + x4)
	fmt.Println(height)
	fmt.Println(result)
	fmt.Println(numberOne, numberTwo, yourName)
}