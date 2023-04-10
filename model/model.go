package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string
	Article string
	Email   string
	Done    string
}

type User struct {
	Name     string
	Email    string
	Password string
}
