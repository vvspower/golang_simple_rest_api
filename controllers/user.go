package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MustafaAP/ProjectK-backend-Go/controllers/helper"
	"github.com/MustafaAP/ProjectK-backend-Go/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://vvspower:lenovo123@cluster0.9ckcd.mongodb.net/?retryWrites=true&w=majority"
const dbName = "ProjectK"
const collectionName = "user"

var collection *mongo.Collection

var mySecretKey = []byte("$sussybaka")

// initialization

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("MongoDB Connection Active")
}

// !helpers

func userNameExists(username string) bool {
	var check model.User
	var exists bool = false
	filter := bson.M{"username": username}
	err := collection.FindOne(context.TODO(), filter).Decode(&check)
	if err == nil {
		exists = true
	}
	return exists
}

func emailExists(email string) bool {
	var check model.User
	var exists bool = false
	filter := bson.M{"email": email}
	err := collection.FindOne(context.TODO(), filter).Decode(&check)
	if err == nil {
		exists = true
	}
	return exists

}

//! controllers

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaton/json")
	var response model.Response
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	email := emailExists(user.Email)
	username := userNameExists(user.Username)

	if email {
		response.Response = "Email already exists"
		response.Success = false
		json.NewEncoder(w).Encode(response)
	} else if username {
		response.Success = false
		response.Response = "Username already exists"
		json.NewEncoder(w).Encode(response)

	} else {
		hashedPassword := helper.HashPass(user.Password)
		user.Password = hashedPassword

		result, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			log.Fatal(err)
		}

		id := result.InsertedID.(primitive.ObjectID).Hex()

		tokenStr := helper.GenerateJWT(id)
		response.Success = true
		response.Response = tokenStr
		json.NewEncoder(w).Encode(response)
	}
}

// bismillah
