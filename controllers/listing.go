package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MustafaAP/ProjectK-backend-Go/model"
	"github.com/gorilla/mux"

	"github.com/MustafaAP/ProjectK-backend-Go/controllers/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection2 *mongo.Collection

func init() {
	collection2 = helper.InitializeDB("listings")
}

func CreateListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	at := r.Header.Get("auth-token")
	data, _ := helper.ExtractClaims(at)

	id, _ := primitive.ObjectIDFromHex(fmt.Sprint(data["userid"]))
	userid := id.Hex()

	var listing model.Listing
	json.NewDecoder(r.Body).Decode(&listing)
	listing.UserID = userid

	result, err := collection2.InsertOne(context.Background(), listing)
	if err != nil {
		log.Fatal(err)
	}

	inserted, _ := result.InsertedID.(primitive.ObjectID)
	listingid := inserted.Hex()
	response := SendResponse(listingid, true)

	json.NewEncoder(w).Encode(response)
}

func GetUserListings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	params := mux.Vars(r)
	username := params["user"]
	fmt.Println(username)
	filter := bson.M{"username": username}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

	id := user.ID.Hex()
	fmt.Println(id)
	filter = bson.M{"userid": id}
	cursor, err := collection2.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	var listings []primitive.M

	for cursor.Next(context.Background()) {
		var listing bson.M
		err := cursor.Decode(&listing)
		if err != nil {
			log.Fatal(err)
		}

		listings = append(listings, listing)
	}

	json.NewEncoder(w).Encode(listings)

}

func GetAllListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cursor, err := collection2.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var listings []primitive.M

	for cursor.Next(context.Background()) {
		var listing bson.M
		err := cursor.Decode(&listing)
		if err != nil {
			log.Fatal(err)
		}

		listings = append(listings, listing)
	}

	json.NewEncoder(w).Encode(listings)

}

func GetOneListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var listing model.Listing

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(fmt.Sprint(params["id"]))

	fmt.Println(id)

	filter := bson.M{"_id": id}

	err := collection2.FindOne(context.TODO(), filter).Decode(&listing)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(listing.Price)
	json.NewEncoder(w).Encode(listing)
}

func UpdateListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var listing model.Listing
	json.NewDecoder(r.Body).Decode(&listing)

	params := mux.Vars(r)
	id := params["id"]

	at := r.Header.Get("auth-token")
	data, _ := helper.ExtractClaims(at)

	userid := data["userid"]

	listingid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"userid": userid, "_id": listingid}
	updated := bson.M{"$set": bson.M{"image": listing.Image, "description": listing.Description, "price": listing.Price}}

	// TODO : update and adjust to new listing method
	_, err := collection2.UpdateOne(context.Background(), filter, updated)
	if err != nil {
		response := SendResponse("Not Found", false)
		w.WriteHeader(404)
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	} else {
		response := SendResponse("List Updated", true)
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}

}

func DeleteListing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	at := r.Header.Get("auth-token")
	data, _ := helper.ExtractClaims(at)

	params := mux.Vars(r)
	id := params["id"]

	userid := data["userid"]
	listingid, _ := primitive.ObjectIDFromHex(id)
	fmt.Println(userid, listingid)

	filter := bson.M{"userid": userid, "_id": listingid}

	deleted, err := collection2.DeleteOne(context.TODO(), filter)

	if err != nil {
		// log.Fatal(err)
		fmt.Println(deleted)
		response := SendResponse("Error occured", false)
		json.NewEncoder(w).Encode(response)
	} else {
		response := SendResponse("Deleted Sucessfully", true)
		json.NewEncoder(w).Encode(response)
	}

}

// TODO: Clean up the Find() code and put it in helper function
// TODO: Clean up error messages
