package service

import (
	"github.com/paulo-fabiano/api-produtos/database"
	"github.com/paulo-fabiano/api-produtos/model"
	"github.com/paulo-fabiano/api-produtos/repository"
)

type ProdutoService struct {
	pr repository.ProdutoRepository
}

func NewProdutoService(pr repository.ProdutoRepository) ProdutoService {
	return ProdutoService{
		pr: pr,
	}
}

func (ps ProdutoService)BuscarProdutosService() []model.Produto{

	listaProdutos, err := database.BuscarProdutos()
	if err != nil {
		panic(err)
	}
	return listaProdutos
}