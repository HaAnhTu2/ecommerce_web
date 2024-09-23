package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	OrderID   primitive.ObjectID `bson:"order_id"`
	ProductID primitive.ObjectID `bson:"product_id"`
	Quantity  int                `bson:"quantity"`
	Price     float64            `bson:"price"`
}
