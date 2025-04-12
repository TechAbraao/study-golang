package main

import "fmt"

func soma(x int, y int) (int) {
	return x + y
}
func sub(x int, y int) (int) {
	return x - y
}
func informations(name string, lastname string) (string, string) {
	return name, lastname
}

func main() {
	var x int = soma(10, 10)
	var y = soma(20, 10)
	var z = sub(40, 1000)
	var name1, name2 string = informations("Abraão", "santos")

	fmt.Printf("\n Resultado: %v\n\n", x)
	fmt.Printf("\n Resultado: %v\n\n", y)
	fmt.Printf("\n Resultado: %v\n\n", z)
	fmt.Printf("\n Resultado: %q e %q\n\n", name1, name2)
}
