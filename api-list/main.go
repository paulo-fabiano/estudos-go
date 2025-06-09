package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Tarefa struct {
	ID int `id`
	Nome string `nome`
	Descricao string `descricao`
}

var tarefas []Tarefa

func main() {

	fmt.Println("##### API Lista de Tarefas #####")
	fmt.Println("##### API Rodando na porta 8000")
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/tasks", handleFindAllTasks)
	http.HandleFunc("/tasks/add", handleAddNewTask)
	http.HandleFunc("/tasks/{id}", handleUpdateTask)
	http.ListenAndServe(":8000", nil)

}

func handlerRoot(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bem vindo a API de Lista de Tarefas"))

}

func handleFindAllTasks(w http.ResponseWriter, r *http.Request){

	// Precisamos primeiro converter a responsta para JSON antes de enviar
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tarefas)

}

func handleAddNewTask(w http.ResponseWriter, r *http.Request) {

	var tarefa Tarefa
	err := json.NewDecoder(r.Body).Decode(&tarefa)	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro"+ err.Error()))
		return 
	}
	tarefas = append(tarefas, tarefa)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tarefa) // Mostrando a tarefa adicioada
	// w.Write([]byte("Tarefa adicionada com sucesso!"))

}

func handleUpdateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var tarefa Tarefa
	err := json.NewDecoder(r.Body).Decode(&tarefa)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro"+ err.Error()))
		return 
	}
	
	for i := range tarefas{
		if tarefas[i].ID == tarefa.ID {
			tarefas[i].Nome = tarefa.Nome
			tarefas[i].Descricao = tarefa.Descricao
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Atualizado com sucesso!"))
			return
		} 
	}
	w.WriteHeader(http.StatusBadRequest)

}