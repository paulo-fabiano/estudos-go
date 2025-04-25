package main

import (

	"fmt"
	"time"

)

func main() {

	// As Goroutines são funções que rodam em paralelo com outras funções.
	// Caracteristicas:
		// - São leves
		// - São gerenciadas pelo Go


	// Exemplo errado:
	go imprimir("goroutine") // goroutine começa, mas é assíncrona
	fmt.Println("Todas as gotoroutines terminaram") // roda imediatamente

	// Exemplo certo:
	go imprimir("goroutine")
	imprimir("normal")
}

func imprimir(mensagem string) {

	for i := 0; i < 10; i++ {
		fmt.Println(mensagem, i)
		time.Sleep(2 * time.Second)
	}

}