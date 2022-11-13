package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	Id          primitive.ObjectID `bson:"_id"`
	Tittle      string             `bson:"tittle"`
	Seller      string             `bson:"seller"`
	Price       float64            `bson:"price"`
	Currency    string             `bson:"currency"`
	Pictures    []string           `bson:"pictures"`
	Description string             `bson:"description"`
	State       string             `bson:"state"`
	City        string             `bson:"city"`
	Street      string             `bson:"street"`
	Number      int                `bson:"number"`
}
