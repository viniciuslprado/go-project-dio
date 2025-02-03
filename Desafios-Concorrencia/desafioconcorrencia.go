// Goroutines: é uma função capaz de executar de modo concorrente com outras funções.
// Canal: modo de duas goroutines se comunicarem e sincronizarem sua execução.
// Select funciona como um Switch para canais. Select permite que você aguarde operações de vários canais.

package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"

	}
}

func pongar(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)

	}
}
func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)
	go pongar(c)

	var entrada string
	fmt.Scanln(&entrada)
}
