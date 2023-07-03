package main

import (
	"fmt"
)

func counter(values int, out chan<- int) {
	fmt.Println("Entrou Counter")
	for i := 1; i <= values; i++ {
		fmt.Printf("valor: %v\n", i)
		out <- i
	}
	close(out)
}

func makeSquares(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Printf("quandrado: %v\n", v)
	}

}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(10, naturals)
	go makeSquares(squares, naturals)
	printer(squares)
	fmt.Println("Fim do programa")
}
