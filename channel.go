package main

import (
	"fmt"
)

func enviarValores(c chan<- int) {
	for i := 1; i <= 5; i++ {
		c <- i // Envia valores para o channel
	}
	close(c) // Fecha o channel após o envio dos valores
}

func receberValores(c <-chan int) {
	for valor := range c {
		fmt.Println(valor) // Recebe e imprime os valores do channel
	}
}

func main() {
	c := make(chan int) // Cria um channel para valores inteiros

	go enviarValores(c) // Chama a função em uma goroutine
	receberValores(c)   // Chama a função na goroutine principal

	fmt.Println("Fim do programa")
}
