package main;

import "fmt"

func main() {

	fmt.Println("Olá, mundo!")
	
	fmt.Println("----- Chamando a função: variaveis() -----")
	variaveis()
	lacoFor()
	condicionais()
}

func variaveis() {
	
	/*
		Podemos declarar variáveis definindo o seu tipo e também fazendo inferências de tipo.
	*/

	var n1 int = 10
	fmt.Println(n1)

	n2 := 10
	fmt.Println(n2)

	// Podemos declarar múltiplas variáveis
	var n3, n4 int = 10, 20
	fmt.Println(n3, n4)


	/*
		Podemos organizar as variáveis em um bloco var()

		Importante notar que não podemos fazer inferência de tipo dentro de um bloco var()
	*/
	var (
		nome string = "Paulo"
		idade int = 25
	)
	fmt.Println(nome, idade)

}

func lacoFor() {

	/*
		O GO só possui o laço for(). Isso é uma vantagem (pelo menos é o que eu acho)
	
		Importante notar que se formos inicializar a variável no for precisamos fazer inferência de tipo,
		se colocar i int = 0 dará erro.
	*/

	fmt.Println("----- Executando um FOR -----")	
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(i)
		i++
	}
	fmt.Println("----- Fim do FOR -----")
	
}

func condicionais() {
	
	idade := 25

	fmt.Println("----- Executando os Condicionais -----")
	if idade > 18 {
		fmt.Println("Menor de idade!")
	} else if idade < 100 {
		fmt.Println("Rapaz, tá quase chegando a hora hein!")
	} else {
		fmt.Println("Maior de idade!")
	}
	fmt.Println("----- Fim dos condicionais -----")
}
