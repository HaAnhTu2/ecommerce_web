package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description"`
	Stock       int                `json:"stock,omitempty" bson:"stock"`
	Price       float64            `json:"price,omitempty" bson:"price"`
	Category    string             `json:"category,omitempty" bson:"category"`
}
