package main

import "fmt"

// Definindo uma função que soma dois valores inteiros.
func sum(x int, y int) int {
	return x + y
}

// Também tem essa forma de fazer uma função.
func summ(x, y int) int {
	return x + y
}


func main() {
	value := sum(10, 50)
	fmt.Println(value)
}