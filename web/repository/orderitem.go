package repository

import (
	"context"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderItemInterface interface {
	FindByID(c context.Context, id string) (model.OrderItem, error)
	Create(c context.Context, orderItem model.OrderItem) (model.OrderItem, error)
	Update(c context.Context, orderItem model.OrderItem) (model.OrderItem, error)
	Delete(c context.Context, id string) error
}

type OrderItemRepo struct {
	DB *mongo.Database
}

func NewOrderItemRepo(db *mongo.Database) OrderItemInterface {
	return &OrderItemRepo{DB: db}
}

func (oi *OrderItemRepo) FindByID(c context.Context, id string) (model.OrderItem, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.OrderItem{}, err
	}
	var order_item model.OrderItem
	err = oi.DB.Collection("order_items").FindOne(c, bson.M{"_id": ID}).Decode(&order_item)
	if err != nil {
		return model.OrderItem{}, err
	}
	return model.OrderItem{}, nil
}

func (oi *OrderItemRepo) Create(c context.Context, orderItem model.OrderItem) (model.OrderItem, error) {
	result, err := oi.DB.Collection("order_items").InsertOne(c, orderItem)
	if err != nil {
		return model.OrderItem{}, err
	}
	orderItem.ID = result.InsertedID.(primitive.ObjectID)
	return orderItem, nil
}

func (oi *OrderItemRepo) Update(c context.Context, orderItem model.OrderItem) (model.OrderItem, error) {

	result, err := oi.DB.Collection("order_items").UpdateOne(c, bson.M{"_id": orderItem.ID}, bson.M{
		"$set": bson.M{
			"quantity": orderItem.Quantity,
			"price":    orderItem.Price,
		},
	})
	if err != nil {
		return model.OrderItem{}, err
	}

	if result.MatchedCount == 0 {
		return model.OrderItem{}, mongo.ErrNoDocuments
	}

	return model.OrderItem{}, nil
}

func (oi *OrderItemRepo) Delete(c context.Context, id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := oi.DB.Collection("order_items").DeleteOne(c, bson.M{"_id": ID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
