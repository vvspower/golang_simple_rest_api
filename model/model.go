package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Firstname string             `json:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty"`
	Email     string             `json:"email,omitempty"`
	Username  string             `json:"username,omitempty"`
	Password  string             `json:"password,omitempty"`
	Contact   string             `json:"contact,omitempty"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Response string `json:"response,omitempty"`
	Success  bool   `json:"success"`
}
