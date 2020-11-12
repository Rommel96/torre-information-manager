package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"-" bson:"_id"`
	Email    string             `json:"email"`
	Password string             `json:"-"`
}
