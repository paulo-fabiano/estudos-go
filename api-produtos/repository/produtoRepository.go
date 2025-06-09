package repository

import (
	"database/sql"

	"github.com/paulo-fabiano/api-produtos/database"
)

type ProdutoRepository struct {
	conexao *sql.DB
}

func NewProdutoRepository() ProdutoRepository {
	return ProdutoRepository{
		conexao: database.Conn,
	}
}