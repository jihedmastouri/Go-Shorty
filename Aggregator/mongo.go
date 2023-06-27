package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveToDB(data interface{}) error {
	collection, err := connectMongo()
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func connectMongo() (*mongo.Collection, error) {
	host := "cluster0.2qj8o.mongodb.net/?w=majority"
	username := "reader"
	psswd := "DWldoNa8losWte27"

	connString := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s",
		host,
		username,
		psswd,
	)
	log.Print(connString)

	clientOptions := options.Client().ApplyURI(connString)
	ctx := context.Background()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	collection := client.Database("shortyio").Collection("blocks")
	return collection, nil
}
