package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/MustafaAP/ProjectK-backend-Go/controllers/helper"
	"github.com/MustafaAP/ProjectK-backend-Go/model"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var collectionFriendReq *mongo.Collection
var collectionFriends *mongo.Collection

func init() {
	collection = helper.InitializeDB("user")
	collectionFriendReq = helper.InitializeDB("friend-req")
	collectionFriends = helper.InitializeDB("friends")
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
	return success
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
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
	w.Header().Set("Content-Type", "application/json")

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

		user.IDCreated = false

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
	w.Header().Set("Content-Type", "application/json")
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
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}
}

func GetUserWithAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserNoAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	params := mux.Vars(r)
	username := params["username"]

	filter := bson.M{"username": username}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = "-"
	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// if u can make another file for this section then make it

//  FRIENDS SECTION

func FriendReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var friendReq model.FriendRequests

	at := r.Header.Get("auth-token")
	data, _ := helper.ExtractClaims(at)
	client := fmt.Sprintf("%s", data["userid"]) // client which sends friend requests

	params := mux.Vars(r)
	id := params["id"] // user which is getting the friend

	filter := bson.M{"to": id, "from": client}

	err := collectionFriendReq.FindOne(context.TODO(), filter).Decode(&friendReq)
	// no fr means erroris not nil

	fmt.Println("hi")
	fmt.Println(err)
	if err == nil {
		response := SendResponse("Friend request already  sent", false)
		json.NewEncoder(w).Encode(response)

	} else {
		friendReq.To = id
		friendReq.From = client

		_, err := collectionFriendReq.InsertOne(context.Background(), friendReq)
		if err != nil {
			log.Fatal(err)
		}
		response := SendResponse("Friend request sent", true)
		json.NewEncoder(w).Encode(response)
	}
}

func GetFriendReqs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	at := r.Header.Get("auth-token")
	data, _ := helper.ExtractClaims(at)
	userid := fmt.Sprintf("%s", data["userid"])

	fmt.Println(fmt.Sprintf("%s", data["userid"]))

	filter := bson.M{"to": userid}

	cursor, err := collectionFriendReq.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var friendReqs []primitive.M

	for cursor.Next(context.Background()) {
		var friendReq bson.M
		err := cursor.Decode(&friendReq)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(friendReq)
		friendReqs = append(friendReqs, friendReq)
	}

	json.NewEncoder(w).Encode(friendReqs)
}
func AcceptFriendReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var friends model.Friends

	json.NewDecoder(r.Body).Decode(&friends)

	rand.Seed(time.Now().UnixNano())
	fmt.Println(randSeq(5))

	friends.Channel = randSeq(5)

	filter := bson.M{"to": friends.UserOne, "from": friends.UserTwo}
	_, err := collectionFriendReq.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	_, error := collectionFriends.InsertOne(context.Background(), friends)
	if error != nil {
		log.Fatal(error)
	}
	response := SendResponse("You are now Friends", true)
	json.NewEncoder(w).Encode(response)
}

func DeleteFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	objectid, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objectid}
	_, err := collectionFriends.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	response := SendResponse("Friend removed", true)
	json.NewEncoder(w).Encode(response)
}

type JSONString string

func (j JSONString) MarshalJSON() ([]byte, error) {
	return []byte(j), nil
}

func GetFriends(w http.ResponseWriter, r *http.Request) {

	at := r.Header.Get("auth-token")
	data, _ := helper.ExtractClaims(at)
	userid := fmt.Sprintf("%s", data["userid"])

	filter := bson.M{"userone": userid}
	cursor, _ := collectionFriends.Find(context.Background(), filter)

	var friends1 []primitive.M

	for cursor.Next(context.Background()) {
		var friend1 bson.M
		err := cursor.Decode(&friend1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(friend1)
		friends1 = append(friends1, friend1)
	}

	filter2 := bson.M{"usertwo": userid}
	cursor2, _ := collectionFriends.Find(context.Background(), filter2)

	var friends2 []primitive.M

	for cursor2.Next(context.Background()) {
		var friend2 bson.M
		err := cursor2.Decode(&friend2)
		if err != nil {
			log.Fatal(err)
		}

		friends2 = append(friends2, friend2)
	}

	friends1 = append(friends1, friends2...)

	json.NewEncoder(w).Encode(friends1) //sent as a string. converted to json in the front end
}

// TODO : find freinds of user and return
