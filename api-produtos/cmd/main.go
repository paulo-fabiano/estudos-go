package main

import (
	"fmt"
	"net/http"

	"github.com/paulo-fabiano/api-produtos/controller"
	"github.com/paulo-fabiano/api-produtos/database"
	"github.com/paulo-fabiano/api-produtos/repository"
	"github.com/paulo-fabiano/api-produtos/service"
)

func main() {

	err := database.InitDB()
	if (err != nil) {
		fmt.Printf("Erro: "+err.Error())
	}

	ProdutoRepository := repository.NewProdutoRepository()
	ProdutoService := service.NewProdutoService(ProdutoRepository)
	ProdutoController := controller.NewProdutoController(ProdutoService)
	fmt.Println(ProdutoController)

	// HealthCheck API
	http.HandleFunc("/ping", controller.HandlePong)
	http.HandleFunc("/produtos", ProdutoController.BuscarProdutosHandler)
	http.ListenAndServe(":8080", nil)

}