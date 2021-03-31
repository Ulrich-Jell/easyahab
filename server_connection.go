package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var PartiesCollection = mongo.Collection{}
//var CandidatesCollection = mongo.Collection{}
var Client = mongo.Client{}

var Lummerland = 5

var EasyAhabDatabase = Client.Database("easyahab")
var PartiesCollection = EasyAhabDatabase.Collection("parties")
var CandidatesCollection = EasyAhabDatabase.Collection("episodes")

func ConnectToServer() {
	Client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckServerConnection() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func Disconnect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	Client.Disconnect(ctx)
	fmt.Println("Connection to MongoDB closed.")
}
