package main

import (
	"fmt"
	"time"
)

func exibirMensagem(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go exibirMensagem("Olá")
	exibirMensagem("Mundo")
}
