package core

import (
	"fmt"
	"bufio"
	"os"
)

func BufioInGo() {
	fmt.Println("Digite seu nome: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		texto := scanner.Text()
		fmt.Println(texto)
	}
}