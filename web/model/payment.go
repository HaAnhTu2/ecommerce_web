package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	OrderID       primitive.ObjectID `bson:"order_id"`
	PaymentMethod string             `bson:"payment_method"`
	Status        string             `bosn:"status"`
}
