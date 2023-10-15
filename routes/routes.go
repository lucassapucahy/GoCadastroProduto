package routes

import (
	"GoCadastroProduto/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/NovoProduto", controllers.NovoProduto)
	http.HandleFunc("/Inserir", controllers.Inserir)
	http.HandleFunc("/deletar/", controllers.Deletar)
	http.HandleFunc("/Editar", controllers.Editar)
	http.HandleFunc("/AlterarProduto/", controllers.AlterarProduto)

}
