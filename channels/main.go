package main

import (
	"fmt"
	"time"
)

func main() {

	// As Channels são canais de comunicação entre goroutines.

	// Para criar um channel
	// channel := make(chan int)
	// Para enviar dados para o channel
	// channel <- 10
	// Para receber dados do channel
	// valor := <-channel
	// fmt.Println(valor) 

	type Dominio struct {
		Nome string
		Valor int
		Verificado bool
	}

	// Criando um channel
	resultadoRequisicao := make(chan bool)

	site := Dominio{
		Nome: "google.com",
		Valor: 10,
		Verificado: false,
	}

	for i := 0; i < 1; i++ {
		go verificarDominio(site.Nome, resultadoRequisicao) // Chamando a função verificarDominio() em uma goroutine
	}

	for i := 0; i < 1; i++ {
		resultado := <-resultadoRequisicao // Recebendo o resultado da requisição
		if resultado {
			fmt.Println("Domínio verificado com sucesso!")
		} else {
			fmt.Println("Erro ao verificar domínio.")
		}
	}


}

func verificarDominio(dominio string, resultadoRequisicao chan bool) {
	// Simulando uma requisição
	time.Sleep(2 * time.Second)
	fmt.Println("Verificando domínio:", dominio)
	resultadoRequisicao <- true // Enviando o resultado da requisição para o channel
}
