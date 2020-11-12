package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Job struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	JobId     string             `json:"jobId"`
	Objective string             `json:"objective"`
	StableOn  string             `json:"stableOn"`
	Deadline  string             `json:"deadline"`
	Status    string             `json:"status"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
}
