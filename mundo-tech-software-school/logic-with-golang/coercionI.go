package main

import "fmt"

func main() {

	var a int = 46
	var b float64 = 6.4

	var c int = int(b)

	fmt.Println(a, b, c)

	var d float64 = float64(c)
	fmt.Println(d)
}