package main

import "fmt"

func main() {
	/*
		fmt.Println("Bem-vindo ao sistema de consultar nomes e obter suas idades.")
		ages := map[string]int{
			"Abraão":    23,
			"José":      35,
			"Anthony":   28,
			"Lucas":     19,
			"Ana":       30,
			"Clara":     22,
			"João":      40,
			"Valentina": 25,
			"Yuki":      27,
			"Zara":      24,
			"Pedro":     33,
			"Igor":      31,
			"Fatima":    29,
		}

		searchName := "Abraão"

		if age, ok := ages[searchName]; ok {
			fmt.Printf("ACHOU! %s tem %d anos.\n", searchName, age)
		} else {
			fmt.Println("NÃO ACHOU!")
		}
	*/

	valueOne := "Abraão"
	valueTwo := valueOne

	fmt.Println(&valueOne)
	fmt.Println(&valueTwo)

	ptrToValueOne := &valueOne
	*ptrToValueOne = "José"
	fmt.Println(*ptrToValueOne)
	fmt.Println(valueTwo)
}
