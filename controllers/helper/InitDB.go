package helper

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeDB(colName string) *mongo.Collection {
	const connectionString = "mongodb+srv://vvspower:lenovo123@cluster0.9ckcd.mongodb.net/?retryWrites=true&w=majority"
	const dbName = "ProjectK"
	collectionName := colName

	var collection *mongo.Collection
	// var mySecretKey = []byte("$sussybaka")

	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("MongoDB Connection Active")

	return collection

}
