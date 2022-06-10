package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MustafaAP/ProjectK-backend-Go/controllers/helper"
	"github.com/MustafaAP/ProjectK-backend-Go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

var collectionIGDetails *mongo.Collection

func init() {
	collectionIGDetails = helper.InitializeDB("InGameDetails")
}

func getUserID(r *http.Request) interface{} {
	at := r.Header.Get("auth-token")
	data, _ := helper.ExtractClaims(at)
	id := data["userid"]
	return id
}

func InitGameDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var IGDetails model.InGameDetails

	json.NewDecoder(r.Body).Decode(&IGDetails)
	id := getUserID(r)

	IGDetails.UserID = fmt.Sprint(id)

	_, err := collectionIGDetails.InsertOne(context.Background(), IGDetails)
	if err != nil {
		log.Fatal(err)
	}

	objectId, _ := primitive.ObjectIDFromHex(fmt.Sprint(id))

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"idcreated": true}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}
	response := SendResponse("Details Updated", true)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func UpdateIGD(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var IGDetails model.InGameDetails

	json.NewDecoder(r.Body).Decode(&IGDetails)
	id := getUserID(r)

	filter := bson.M{"userid": fmt.Sprint(id)}
	update := bson.M{"$set": bson.M{"rank": IGDetails.Rank, "time": IGDetails.Time, "role": IGDetails.Role, "discord": IGDetails.Discord}}
	_, err := collectionIGDetails.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		w.WriteHeader(500)
	}
	response := SendResponse("Details Updated", true)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(500)
	}
}
