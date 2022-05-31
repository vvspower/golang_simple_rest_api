package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// user

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Firstname string             `json:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty"`
	Email     string             `json:"email,omitempty"`
	Username  string             `json:"username,omitempty"`
	Password  string             `json:"password,omitempty"`
	Image     string             `json:"image,omitempty"`
	Contact   string             `json:"contact,omitempty"`
}

type UpdateUser struct {
	Image   string `json:"image"`
	Contact string `json:"contact"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Response string `json:"response,omitempty"`
	Success  bool   `json:"success"`
}

//  Listings

type Listing struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID      string             `json:"userid"`
	Image       string             `json:"image"`
	Item        string             `json:"item"`
	Description string             `json:"description"`
	Price       string             `json:"price"`
	Rating      string             `json:"rating"`
}
