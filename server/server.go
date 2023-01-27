package main

import (
	"fmt"
	"net/http"
	po "pokedex"
	pos "pokedex/struct"
	"text/template"
)

var Pokemons []pos.Pokemon

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", HandleIndex)
	fs := http.FileServer(http.Dir("./static"))
	Pokemons = po.GetDatas()
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(w, Pokemons)
	return
}
