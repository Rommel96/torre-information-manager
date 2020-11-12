package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rommel96/torre-information-manager/backend/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// create UniqueIndex option
var userIndexModel = mongo.IndexModel{
	Keys: bson.M{
		"email": 1,
	},
	Options: options.Index().SetUnique(true),
}

func RunConfig() {
	loadEnv()
	repository.DBConnect()
	//initMigration() //just run once time
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
}
