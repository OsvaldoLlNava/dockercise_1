package Dockercise1APIDatabase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/OsvaldoLlNava/dockercise1API/Dockercise1APIModel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() (context.Context, *mongo.Client, context.CancelFunc) {
	uri := "mongodb://mongo_compose:27017/"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	credential := options.Credential{
		Username: "quesitoUser",
		Password: "SecretCheese1234",
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(credential))
	if err != nil {
		fmt.Println(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	return ctx, client, cancel
}

func Disconnect(ctx context.Context, client *mongo.Client, cancel context.CancelFunc) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
	cancel()
}

func ObtenerPersonas() []Dockercise1APIModel.Person {

	ctx, client, cancel := Connect()
	defer Disconnect(ctx, client, cancel)

	db := client.Database("dockercise1")
	collection := db.Collection("Personas")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	var p []Dockercise1APIModel.Person

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &p)
	if err != nil {
		panic(err)
	}

	return p

}

func ObtenerUnaPersona(id string) Dockercise1APIModel.Person {

	ctx, client, cancel := Connect()
	defer Disconnect(ctx, client, cancel)

	var p Dockercise1APIModel.Person

	numero, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		panic(err)
	}

	db := client.Database("dockercise1")
	collection := db.Collection("Personas")
	err = collection.FindOne(ctx, bson.D{{Key: "id", Value: numero}}).Decode(&p)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return p
		} else {
			panic(err)
		}
	}

	return p

}
