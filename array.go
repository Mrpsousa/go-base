package main

import "fmt"

func main() {
	// Declaração e inicialização de um array
	var arr [5]int

	//sem tamanho
	idades := [...]int{20, 30, 40, 50}
	nomes := [3]string{"João", "Maria", "Pedro"}
	// Inserir elementos no array
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	// Imprimir o array
	fmt.Println("Array original:", arr)

	// Acessar um elemento do array
	value := arr[2]
	fmt.Println("Elemento do índice 2:", value)

	// Modificar um elemento do array
	arr[3] = 45
	fmt.Println("Array modificado:", arr)

	// Obter o tamanho do array
	size := len(arr)
	fmt.Println("Tamanho do array:", size)

	// Iterar sobre os elementos do array
	fmt.Println("Elementos do array:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Verificar se o array está vazio
	isEmpty := len(arr) == 0
	fmt.Println("Array está vazio:", isEmpty)
}
