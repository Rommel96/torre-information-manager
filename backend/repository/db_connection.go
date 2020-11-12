package repository

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbName string
	uri    string
	Client *mongo.Client
	err    error
)

func DBConnect() {
	formatUri()
	clientOptions := options.Client().ApplyURI(uri)
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB Connected...")
}

func formatUri() {
	dbName = os.Getenv("DB_NAME")
	uri = "mongodb+srv://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@cluster0.e3ofv.mongodb.net/" + dbName + "?retryWrites=true&w=majority"
}
