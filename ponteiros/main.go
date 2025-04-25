package main

import "fmt"

func main() {

	// Ponteiro é uma variável que armazena o endereço de memório de outra variável.
	// *T declara um ponteiro para o tipo T.
	// &T retorna o endereço de memória da variável T.

	// Exemplo:
	var a int = 10
	var b *int = &a 

	fmt.Println("Valor de a:", a) // 10
	fmt.Println("Endereço de a:", &a) // Endereço de memória de a
	fmt.Println("Valor de b: ", b) // Endereço de b, que é o endereço de memória de a
	fmt.Println("Valor de b: ", *b) // 10

	// Com a função dobrar()
	var numero int = 5
	dobrar(&numero)
	fmt.Println("Valor de numero:", numero) 
}

func dobrar(valor *int) {
	*valor = *valor * 2
}