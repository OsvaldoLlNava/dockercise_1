package Dockercise1Database

import (
	"context"
	"fmt"
	"time"

	"github.com/OsvaldoLlNava/dockercise1/Dockercise1Model"
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

func InsertarDocumento(ctx context.Context, coleccion *mongo.Collection, p Dockercise1Model.Person) {
	_, err := coleccion.InsertOne(ctx, bson.D{
		{Key: "id", Value: p.Id},
		{Key: "first_name", Value: p.First_Name},
		{Key: "last_name", Value: p.Last_Name},
		{Key: "company", Value: p.Company},
		{Key: "email", Value: p.Email},
		{Key: "ip_address", Value: p.Ip_Address},
		{Key: "phone_number", Value: p.Phone_Number},
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Insertado correcto")
	}
}
