package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserName string             `bson:"username"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
	Address  string             `bson:"address"`
	Phone    string             `bson:"phone"`
	Role     string             `bson:"role"`
}

type LoginRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}
