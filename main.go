package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome, Descricao string
	Preco float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome:"Camiseta", Descricao:"Camiseta de futebol", Preco:49.90, Quantidade:10},
		{Nome:"Calça", Descricao:"Calça de algodão", Preco:79.90, Quantidade:20},
		{Nome:"Bermuda", Descricao:"Bermuda de couro", Preco:59.90, Quantidade:30},
		{Nome:"Tênis", Descricao:"Tênis de couro", Preco:129.90, Quantidade:40},
	}
	
	temp.ExecuteTemplate(w, "Index", produtos)
}