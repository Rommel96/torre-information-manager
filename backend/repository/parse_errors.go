package repository

import "go.mongodb.org/mongo-driver/mongo"

func IsDuplicateKey(err error) bool {
	merr := err.(mongo.WriteException)
	//log.Errorf("Number of errors: %d", len(merr.WriteErrors))
	errCode := merr.WriteErrors[0].Code
	if errCode == 11000 {
		return true
	}
	return false
}
