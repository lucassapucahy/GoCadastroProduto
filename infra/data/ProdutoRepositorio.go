package repositorios

import (
	"GoCadastroProduto/models"

	"database/sql"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	connStr := "user=postgres dbname=gocadastroproduto password=Postgres2022! host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func BuscaProdutos() []models.Produto {
	db := conectaComBancoDeDados()
	defer db.Close()

	result, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := models.Produto{}
	produtos := []models.Produto{}

	for result.Next() {
		err := result.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtos = append(produtos, p)
	}

	return produtos
}
