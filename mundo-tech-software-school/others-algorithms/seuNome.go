package main

import "fmt"

func seuNome(nome string) string {
	if nome == "" {
		return "Nenhum nome definido"
	}
	return "Seu nome é: " + nome
}


func main() {
	fmt.Println(seuNome("Abraão"))
	fmt.Println(seuNome(""))
}
