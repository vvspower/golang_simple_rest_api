package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/MustafaAP/ProjectK-backend-Go/controllers/helper"
	"github.com/MustafaAP/ProjectK-backend-Go/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	collection = helper.InitializeDB("user")
}

// !helpers

func checkEmpty(user model.User) bool {
	var success bool = true
	if user.Firstname == "" {
		success = false
	}
	if user.Lastname == "" {
		success = false
	}
	if user.Username == "" {
		success = false
	}
	if user.Email == "" {
		success = false
	}
	if user.Password == "" {
		success = false
	}
	if user.Contact == "" {
		success = false
	}
	return success
}

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

func SendResponse(res string, success bool) model.Response {
	var response model.Response
	response.Response = res
	response.Success = success
	return response
}

//! controllers

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaton/json")

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	email := emailExists(user.Email)
	username := userNameExists(user.Username)
	if !checkEmpty(user) {
		response := SendResponse("Something went wrong. Please try again", false)
		json.NewEncoder(w).Encode(response)
	} else if email {
		response := SendResponse("Email already exists", false)
		json.NewEncoder(w).Encode(response)
	} else if username {
		response := SendResponse("Username already exists", false)
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
		response := SendResponse(tokenStr, true)

		json.NewEncoder(w).Encode(response)
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaton/json")
	var login model.Login
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"email": login.Email}
	collection.FindOne(context.TODO(), filter).Decode(&user)

	hashedPass := []byte(user.Password)
	originalPass := []byte(login.Password)

	err = bcrypt.CompareHashAndPassword(hashedPass, originalPass)
	if err == nil {
		tokenStr := helper.GenerateJWT(user.ID.Hex())
		response := SendResponse(tokenStr, true)
		json.NewEncoder(w).Encode(response)

	} else {
		response := SendResponse("use correct credentials", false)
		json.NewEncoder(w).Encode(response)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaton/json")
	var user model.User

	at := r.Header.Get("auth-token")
	data, _ := helper.ExtractClaims(at)

	id := data["userid"]

	objectId, _ := primitive.ObjectIDFromHex(fmt.Sprint(id))

	filter := bson.M{"_id": objectId}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = "-"
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicaton/json")
	var update model.UpdateUser

	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Fatal(err)
	}
	at := r.Header.Get("auth-token")
	data, _ := helper.ExtractClaims(at)

	id, _ := primitive.ObjectIDFromHex(fmt.Sprint(data["userid"]))

	filter := bson.M{"_id": id}
	updated := bson.M{"$set": bson.M{"image": update.Image, "contact": update.Contact}}
	_, er := collection.UpdateOne(context.Background(), filter, updated)
	if er != nil {
		log.Fatal(er)
	}

	response := SendResponse("Updated", true)

	json.NewEncoder(w).Encode(response)
}

// bismillah
