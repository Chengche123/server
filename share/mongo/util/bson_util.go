package util

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const FiledID = "_id"

type ObjID struct {
	ID primitive.ObjectID `bson:"_id"`
}

// Set return $set wrapper util function
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

// SetOnInsert return $SetOnInsert wrapper util function
func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}
