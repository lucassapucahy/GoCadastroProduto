package main

import (
	repositorios "GoCadastroProduto/infra/data"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := repositorios.BuscaProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}
