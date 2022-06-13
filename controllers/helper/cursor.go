package helper

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func CursorHelper(cursor *mongo.Cursor) []primitive.M {
	var array []primitive.M
	for cursor.Next(context.Background()) {
		var single bson.M
		err := cursor.Decode(&single)
		if err != nil {
			log.Fatal(err)
		}
		array = append(array, single)
	}
	return array
}
