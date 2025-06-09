package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/paulo-fabiano/api-produtos/model"
)

// type Database struct {
// 	conexao *sql.DB
// }

const (
	host string = "localhost";
	port string = "5432";
	user string = "postgres";
	password string = "12345";
	database string = "estudos_go"
)

var (
	Conn *sql.DB
)

func InitDB() (error) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	conn, err := sql.Open("postgres", psqlInfo)
	if (err != nil) {
		err.Error()
		return err
	}

	err = conn.Ping()
	if (err != nil) {
		fmt.Printf("Erro ao testar conexão no banco de dados: "+err.Error())
		return fmt.Errorf("Erro"+err.Error())
	}
	Conn = conn
	return nil

}

func BuscarProdutos() ([]model.Produto, error) {

	var listaProdutos []model.Produto
	var produto model.Produto

	query := "SELECT * FROM produtos"
	stmt, err := Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}	
	defer rows.Close()
	
	for rows.Next() {
		err = rows.Scan(
			&produto.ID,
			&produto.Nome,
			&produto.Descricao,
			&produto.Quantidade,
			&produto.Preco,
		)
		if err != nil {
			panic(err)
			// return nil, pa
		}
		listaProdutos = append(listaProdutos, produto)
	}

	return listaProdutos, nil
}