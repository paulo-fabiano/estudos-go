package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Um context em Go define o tipo de contexto, que transporta
		- prazos
		- sinais de cancelamento, tipo se a API não responder em x segundos retorne 500

	Quando iniciamos operações em Go é bastante comum criar um context
*/

func main()  {
	
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	bookHotel(ctx)

}

// Quando trabalhamos com context, devemos passar ele em primeiro nos parâmetros por questão de boas práticas
func bookHotel(ctx context.Context) {

	select {
	case <- ctx.Done():
		fmt.Println("Tempo excedido para bookar o quarto!")
	case <- time.After(time.Second * 5):
		fmt.Println("Quarto reservado com sucesso!")
	}

}