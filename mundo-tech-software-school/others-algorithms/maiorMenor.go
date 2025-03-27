package main

import "fmt"

func maiorMenor(operador, numero int) string {
	if operador >= numero {
		return "Operador é maior ou igual ao número escolhido"
	} else if operador <= numero {
		return "Operador é menor ou igual ao número escolhido"
	}
	return "Operador igual ao número escolhido"
}


func main() {
	fmt.Println(maiorMenor(2, 5))
}