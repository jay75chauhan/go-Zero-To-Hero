package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movies  string `json:"movies,omitempty"`
	Watched bool  `json:"watched,omitempty"`
}