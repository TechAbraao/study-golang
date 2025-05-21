package utils
import "fmt"

// Arrays
func ShowArray() [5]int {
	// Isso é um Array no Golang
	var numbers = [5]int{1, 2, 3, 4, 5}

	return numbers
}

// Arrays
func AddingNamesInArray() {
	var names = [5]string{}
	fmt.Println("Digite o nome dos cinco usuários que deseja adicionar no Array")

	// Laço de repetição 'for' no Golang
	for i := 0; i <= 4; i++ {
		fmt.Scanln(&names[i])
		fmt.Println("Nome: ", names[i])
	}

	// Utilizando o conceito de For..In em Arrays
	fmt.Println()
	for i, name := range names {
		fmt.Println("Índice: ", i, " Nome: ", name)
	}

}

// Slices no Golang
func ElasticArrayNumbers() {
	var quantity int
	var addingNumber int

	fmt.Println("Quantos números deseja adicionar na lista?")
	fmt.Scanln(&quantity)

	// Criando um Slice, cujo tamanho é flexível e depende diretamente da variável 'quantity'
	var numbersSlice = make([]int, quantity)

	// Adicionando valores no 'numbersSlice'
	fmt.Println("Agora vamos adicionar os números ", quantity, " números escolhidos!")
	for i := 0; i < quantity; i++ {
		fmt.Scanln(&addingNumber)
		numbersSlice[i] = addingNumber
	}
	fmt.Println("Todos os números: ", numbersSlice)
}