package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

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

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/home.html")
	if err != nil {
		fmt.Fprintf(w, "Bienvenido al Home")
	}

	err = t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func handleRequests() {
	r := chi.NewRouter()
	r.Get("/", homePage)
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
