package repository

import (
	"context"

	"github.com/rommel96/torre-information-manager/backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertJob(Job *models.Job) error {
	_, err := Client.Database(dbName).Collection("jobs").InsertOne(context.TODO(), &Job)
	if err != nil {
		return err
	}
	return nil
}

func DeleteJob(Job *models.Job) error {
	_, err := Client.Database(dbName).Collection("jobs").DeleteOne(context.TODO(), &Job)
	if err != nil {
		return err
	}
	return nil
}

func FindJobById(_id string) (*models.Job, error) {
	if err != nil {
		return nil, err
	}
	var job models.Job
	err = Client.Database(dbName).Collection("users").FindOne(context.TODO(), bson.M{
		"_id": _id,
	}).Decode(&job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}
