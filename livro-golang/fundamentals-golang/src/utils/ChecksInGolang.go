package utils

import "fmt"

func Check() {
	names := []string{"Abraão", "Vitor", "José"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
