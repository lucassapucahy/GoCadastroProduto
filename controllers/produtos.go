package controllers

import (
	repositorios "GoCadastroProduto/infra/data"
	"GoCadastroProduto/models"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := repositorios.BuscaProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}

func NovoProduto(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NovoProduto", nil)
}

func Inserir(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := models.Produto{}

		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)

		if err != nil {
			panic("Erro ao converter preço: " + r.FormValue("preco") + err.Error())
		}

		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))

		if err != nil {
			panic("Erro ao converter quantidade" + r.FormValue("quantidade") + err.Error())
		}

		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")
		produto.Preco = preco
		produto.Quantidade = quantidade

		repositorios.InsereProduto(produto)

		http.Redirect(w, r, "/", 301)
	}
}

func Deletar(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) < 2 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	convResult, convErr := strconv.Atoi(pathParts[2])

	if convErr != nil {
		http.Error(w, r.URL.Path, http.StatusBadRequest)
	}

	repositorios.DeletarProduto(convResult)

	http.Redirect(w, r, "/", 301)
}

func AlterarProduto(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 2 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idConvertido, convErr := strconv.Atoi(pathParts[2])

	if convErr != nil {
		http.Error(w, r.URL.Path, http.StatusBadRequest)
	}

	produto := repositorios.BuscaProduto(idConvertido)

	templates.ExecuteTemplate(w, "EditarProduto", produto)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		produto := models.Produto{}

		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)

		if err != nil {
			panic("Erro ao converter preço: " + r.FormValue("preco") + err.Error())
		}

		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))

		if err != nil {
			panic("Erro ao converter quantidade" + r.FormValue("quantidade") + err.Error())
		}

		id, err := strconv.Atoi(r.FormValue("id"))

		if err != nil {
			panic("Erro ao converter Id" + r.FormValue("id") + err.Error())
		}

		produto.Id = id
		produto.Nome = r.FormValue("nome")
		produto.Descricao = r.FormValue("descricao")
		produto.Preco = preco
		produto.Quantidade = quantidade

		repositorios.AlteraProduto(produto)

		http.Redirect(w, r, "/", 301)
	}
}
