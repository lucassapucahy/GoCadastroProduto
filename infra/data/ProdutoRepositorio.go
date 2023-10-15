package repositorios

import (
	"GoCadastroProduto/models"

	"database/sql"

	_ "github.com/lib/pq"

	"log"
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

	result, err := db.Query("select * from produtos order by Id")

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

func BuscaProduto(Id int) models.Produto {
	db := conectaComBancoDeDados()
	defer db.Close()

	result := db.QueryRow("select * from produtos where Id = $1 ;", Id)

	p := models.Produto{}

	errScan := result.Scan(&p.Id, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)

	switch {
	case errScan == sql.ErrNoRows:
		log.Printf("produto nao encontrado %d\n", Id)
	case errScan != nil:
		log.Fatalf("query error: %v\n", errScan)
	}

	return p
}

func InsereProduto(produto models.Produto) int {
	db := conectaComBancoDeDados()
	defer db.Close()

	id := 0

	stmt, err := db.Prepare("INSERT INTO produtos (nome,descricao,preco,quantidade) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade)

	return id
}

func AlteraProduto(produto models.Produto) int {
	db := conectaComBancoDeDados()
	defer db.Close()

	id := 0

	stmt, err := db.Prepare("UPDATE produtos SET nome = $1, descricao= $2, preco= $3, quantidade = $4 WHERE id = $5")

	if err != nil {
		panic("test" + err.Error())
	}

	stmt.Exec(produto.Nome, produto.Descricao, produto.Preco, produto.Quantidade, produto.Id)

	return id
}

func DeletarProduto(id int) bool {
	db := conectaComBancoDeDados()
	defer db.Close()

	stmt, err := db.Prepare("delete from produtos where Id = $1")

	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Exec(id)

	if err != nil {
		panic(err.Error())
	}

	RowsAffected, _ := result.RowsAffected()

	return RowsAffected == 1

}
