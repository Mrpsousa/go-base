package main

import "fmt"

func main() {
	// Declaração e inicialização de um slice
	slice := []int{1, 2, 3, 4, 5}

	// Imprimir o slice
	fmt.Println("Slice original:", slice)

	// Acessar um elemento do slice
	fmt.Println("Primeiro elemento:", slice[0])

	// Modificar um elemento do slice
	slice[2] = 10
	fmt.Println("Slice modificado:", slice)

	// Obter o comprimento do slice
	fmt.Println("Comprimento do slice:", len(slice))

	// Adicionar elementos ao slice
	slice = append(slice, 6, 7)
	fmt.Println("Slice após adicionar elementos:", slice)

	// Copiar um slice para outro
	clone := make([]int, len(slice))
	copy(clone, slice)
	fmt.Println("Cópia do slice:", clone)

	// Obter uma parte do slice (subslice)
	subslice := slice[1:4]
	fmt.Println("Subslice:", subslice)

	// Remover elementos do slice
	slice = append(slice[:2], slice[4:]...)
	fmt.Println("Slice após remover elementos:", slice)

	// Iterar sobre os elementos do slice
	fmt.Print("Elementos do slice:")
	for _, value := range slice {
		fmt.Print(" ", value)
	}
	fmt.Println()

	// Verificar se um elemento está presente no slice
	element := 3
	found := false
	for _, value := range slice {
		if value == element {
			found = true
			break
		}
	}
	fmt.Println("Elemento", element, "encontrado:", found)
}
