package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/OsvaldoLlNava/dockercise1/Dockercise1Database"
	"github.com/OsvaldoLlNava/dockercise1/Dockercise1Model"
)

func main() {

	ctx, client, cancel := Dockercise1Database.Connect()
	defer Dockercise1Database.Disconnect(ctx, client, cancel)

	// leer el archivo
	content, err := ioutil.ReadFile("./info/people.xml")
	if err != nil {
		panic(err)
	}
	// fmt.Println(content)

	var people Dockercise1Model.People

	err = xml.Unmarshal(content, &people)
	if err != nil {
		panic(err)
	}
	db := client.Database("dockercise1")
	coleccion := db.Collection("Personas")

	for _, v := range people.Personas {
		Dockercise1Database.InsertarDocumento(ctx, coleccion, v)
	}

	fmt.Println("Proceso Finalizado")

}
