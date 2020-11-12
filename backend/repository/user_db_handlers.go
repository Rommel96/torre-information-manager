package repository

import (
	"context"

	"github.com/rommel96/torre-information-manager/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(user models.SignupModel) (interface{}, error) {
	collection := Client.Database(dbName).Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

func FindUserFromLogin(dataLogin models.LoginModel) (*models.User, error) {
	var user models.User
	err := Client.Database(dbName).Collection("users").FindOne(context.TODO(), bson.M{
		"email": dataLogin.Email,
	}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserById(idString string) (*models.User, error) {
	_id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = Client.Database(dbName).Collection("users").FindOne(context.TODO(), bson.M{
		"_id": _id,
	}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// func ChangePassword(_id primitive.ObjectID, newHashedPassword string) (*models.User, error) {
// 	var user models.User
// 	err := Client.Database(dbName).Collection("users").FindOneAndUpdate(context.TODO(), bson.M{
// 		"_id": _id,
// 	}, bson.D{
// 		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "password", Value: newHashedPassword}}},
// 	},
// 	).Decode(&user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func DeleteUserById(idString string) (*models.User, error) {
// 	_id, err := primitive.ObjectIDFromHex(idString)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var user models.User
// 	err = Client.Database(dbName).Collection("users").FindOneAndDelete(context.TODO(), bson.M{
// 		"_id": _id,
// 	}).Decode(&user)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
