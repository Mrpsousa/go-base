package main

import "fmt"


//var m = make(map[string]int)

func main() {
	// Declaração e inicialização de um map
	m := make(map[string]int)

	// Inserir elementos no map
	m["Alice"] = 25
	m["Bob"] = 30
	m["Charlie"] = 35

	// Imprimir o map
	fmt.Println("Map original:", m)

	// Acessar um valor do map
	age := m["Bob"]
	fmt.Println("Idade do Bob:", age)

	// Modificar um valor do map
	m["Alice"] = 26
	fmt.Println("Map modificado:", m)

	// Verificar se uma chave existe no map
	_, exists := m["Charlie"]
	fmt.Println("Chave Charlie existe:", exists)

	// Remover um elemento do map
	delete(m, "Bob")
	fmt.Println("Map após remover Bob:", m)

	// Iterar sobre as chaves e valores do map
	fmt.Println("Chaves e Valores do Map:")
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

	// Obter o número de elementos no map
	fmt.Println("Número de elementos no Map:", len(m))

	// Verificar se o map está vazio
	isEmpty := len(m) == 0
	fmt.Println("Map está vazio:", isEmpty)
}
