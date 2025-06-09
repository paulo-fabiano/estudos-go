package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	var cliente Cliente = Cliente{"Paulo", 26, true}
	imprimirCliente(cliente)

	http.HandleFunc("/ping", handleHealthCheck)
	http.ListenAndServe(":8000", nil)
}

type Cliente struct {
	Nome string
	Idade int8
	Ativo bool
}

func imprimirCliente(cliente Cliente) {
	fmt.Printf("Cliente: %s - %d anos - Ativo: %t", cliente.Nome, cliente.Idade, cliente.Ativo)
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	message := map[string]string {"message" : "pong"}
	json.NewEncoder(w).Encode(message)

}

func somaPares(valores []int) int {
	var soma int = 0
	for _, numero := range valores {
		if numero % 2 == 0 {
			soma +=numero
		}
	}
	return soma
}
