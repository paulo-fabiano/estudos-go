package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {

	// O WaitGroup é o responsável por sincronizar as Goroutines.
	// Ele é um contador que espera todas as Goroutines terminarem antes de continuar a execução do programa.

	// Exemplo:
	var wg sync.WaitGroup // Criando um WaitGroup
	wg.Add(1)
	go imprimir("goroutine", &wg) // Chamando a função imprimir() em uma goroutine
	wg.Wait() // Esperando todas as goroutines terminarem
	fmt.Println("Todas as goroutines terminaram") 
}

func imprimir(mensagem string, wg *sync.WaitGroup) {

	defer wg.Done() // Uma das tarefas terminou, decrementando o contador do WaitGroup

	for i := 0; i < 10; i++ {
		fmt.Println(mensagem, i)
		time.Sleep(1 * time.Second)
	}

}