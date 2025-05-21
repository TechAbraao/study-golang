package utils

import "fmt"

func TriangleAlgorithm() {
	asterisk := ""

	for i := 0; i <= 10; i++ {
		asterisk += "*"
		fmt.Println(asterisk)
	}
}