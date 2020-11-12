package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	Id     string             `json:"jobId" bson:"_id"`
	UserId primitive.ObjectID `json:"userId" bson:"userId"`
}
