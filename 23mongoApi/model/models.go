package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie string `json:"movie,omitempty"`
	Watched bool `json:"watched,omitempty"`
	//The omitempty option on the json and bson tags ensures that the ID field is not included in the JSON or BSON representation of the struct if it is empty or zero.
}