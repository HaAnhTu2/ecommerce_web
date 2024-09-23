package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	TotalAmount float64            `bson:"total_amount"`
	OrderDate   time.Time          `bson:"order_date"`
	Status      string             `bson:"status"`
}
