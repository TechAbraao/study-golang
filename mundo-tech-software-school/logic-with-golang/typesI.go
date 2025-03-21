
/*	
	int // Tipo Inteiro. Existem vários tipos de int, ex: int8, int16, int32, int64 ... 
	// São ints de vários bytes diferentes
	string // Tipo Stirng
	uint // Aceita apena números de 0 acima, ou seja, apenas positivos

	float32 // Por padrão, caso não especifique o byte dele
	float64 // float de 65 bytes

	bool // Podendo ser true ou false

	complex64 // Pra representar os números complexos
	complex128
*/
package main

import "fmt"

func main() {
	var height = 1.34
	var isVerify bool = false

	// O formtato %T imprime o TIPO, enquanto o formato %v imprime o valor atribuído a variável
	fmt.Printf("Tipo: %T  Valor: %v\n", height, height)

	// Aqui já verifica se é true, isto é, (isVerify == true)
	if (isVerify) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}

}	
