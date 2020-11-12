package config

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/rommel96/torre-information-manager/backend/middleware"
	"github.com/rommel96/torre-information-manager/backend/models"
	"github.com/rommel96/torre-information-manager/backend/repository"
	"github.com/rommel96/torre-information-manager/backend/utils"
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

func initMigration() {
	client := repository.Client
	email := "rommel@admin.com"
	passHashed, err := utils.Hash("Abcd1234_")
	if err != nil {
		panic(err)
	}
	//Create User Admin
	userCollection := client.Database("torre-test").Collection("users")
	_, err = userCollection.InsertOne(context.TODO(), models.SignupModel{
		Name:     "Rommel",
		Email:    email,
		Password: string(passHashed),
	})
	if err != nil {
		panic(err)
	}
	_, err = userCollection.Indexes().CreateOne(context.TODO(), userIndexModel)
	if err != nil {
		panic(err)
	}
	token, err := middleware.GenerateToken(email)
	if err != nil {
		panic(err)
	}
	log.Println("TOKEN ADMIN: ", token)
}
