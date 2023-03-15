package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Iniciando servidor na porta 7000")
	err := http.ListenAndServe(":7000", mux)
	log.Fatal(err)
}
