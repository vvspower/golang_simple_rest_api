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
	Followers []string           `json:"followers"`
	Following []string           `json:"following"`
	IDCreated bool               `json:"idcreated"`
}

type FriendRequests struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	To   string             `json:"to"`
	From string             `json:"from"`
}

type Friends struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	UserOne string             `json:"user-one"`
	UserTwo string             `json:"user-two"`
	Channel string             `json:"channel"`
}

type InGameDetails struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  string             `json:"userid"`
	Rank    string             `json:"rank"`
	Time    []string           `json:"time"` // will be like time : ["9:00AM", "6:00PM"] meaning 9am to 6pm
	Role    string             `json:"role"`
	Discord string             `json:"discord"`
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
	Description string             `json:"description"`
	FEmail      bool               `json:"femail"`
	Bans        bool               `json:"bans"`
	Region      string             `json:"region"`
	Skins       []string           `json:"skins"`
	Payment     string             `json:"payment"`
	Price       string             `json:"price"`
}

type UpdateListing struct {
	Image       string `json:"image"`
	Item        string `json:"item"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

// FORUM

type Post struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  string             `json:"userid"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
	Likes   []string           `json:"likes"`
	Tags    []string           `json:"tags"`
}

type Reply struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  string             `json:"userid"`
	PostID  string             `json:"post-id"`
	Content string             `json:"content"`
}
