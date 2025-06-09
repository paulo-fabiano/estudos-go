package controller

import (
	"encoding/json"
	"net/http"

	"github.com/paulo-fabiano/api-produtos/model"
	"github.com/paulo-fabiano/api-produtos/service"
)

type ProdutoController struct{
	ps service.ProdutoService
}

func NewProdutoController(ps service.ProdutoService) ProdutoController{
	return ProdutoController{
		ps: ps,
	}
}

func HandlePong(w http.ResponseWriter, r *http.Request) {

	message := model.Pong{
		"pong",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)

}

func (pc *ProdutoController) BuscarProdutosHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	lista := pc.ps.BuscarProdutosService()
	json.NewEncoder(w).Encode(lista)

}