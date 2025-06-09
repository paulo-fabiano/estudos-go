package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/test", handleHome)
	http.ListenAndServe(":8080", nil)

}

func handleHome(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	log.Printf("Inicio minha request")
	defer log.Println("Finalizando minha request")

	select {
	case <- time.After(time.Second * 5):
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Requisição processada com sucesso"))
	case <- ctx.Done():
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)
	}

}