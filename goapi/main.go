package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OsvaldoLlNava/dockercise1API/Dockercise1APIDatabase"
	"github.com/OsvaldoLlNava/dockercise1API/Dockercise1APIModel"
	"github.com/go-chi/chi/v5"
)

func allResults(w http.ResponseWriter, r *http.Request) {
	p := obtenerTodo()

	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		panic(err)
	}
}

func specificResult(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userId")
	p := obtenerPersona(id)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		panic(err)
	}
}

func handleRequests() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "bienvenido al home")
	})
	r.Get("/people", allResults)
	r.Get("/people/{userId}", specificResult)
	http.ListenAndServe(":7777", r)
}

func obtenerTodo() []Dockercise1APIModel.Person {

	p := Dockercise1APIDatabase.ObtenerPersonas()

	return p
}

func obtenerPersona(i string) Dockercise1APIModel.Person {
	p := Dockercise1APIDatabase.ObtenerUnaPersona(i)
	return p
}

func main() {
	handleRequests()
}
