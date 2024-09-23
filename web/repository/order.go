package repository

import (
	"context"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepoInterface interface {
	FindByID(c context.Context, id string) (model.Order, error)
	Create(c context.Context, order model.Order) (model.Order, error)
	Update(c context.Context, order model.Order) (model.Order, error)
	Delete(c context.Context, id string) error
}
type OrderRepo struct {
	DB *mongo.Database
}

func NewOrderRepo(db *mongo.Database) OrderRepoInterface {
	return &OrderRepo{DB: db}
}

func (o *OrderRepo) FindByID(c context.Context, id string) (model.Order, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Order{}, err
	}
	var order model.Order
	err = o.DB.Collection("order").FindOne(c, bson.M{"_id": ID}).Decode(&order)
	if err != nil {
		return model.Order{}, err
	}
	return model.Order{}, nil
}

func (o *OrderRepo) Create(c context.Context, order model.Order) (model.Order, error) {
	result, err := o.DB.Collection("order").InsertOne(c, order)
	if err != nil {
		return model.Order{}, err
	}
	order.ID = result.InsertedID.(primitive.ObjectID)
	return order, nil
}

func (o *OrderRepo) Update(c context.Context, order model.Order) (model.Order, error) {
	result, err := o.DB.Collection("order").UpdateOne(c, bson.M{"_id": order.ID}, bson.M{
		"$set": bson.M{
			"total_amount": order.TotalAmount,
			"status":       order.Status,
		},
	})
	if err != nil {
		return model.Order{}, err
	}
	if result.MatchedCount == 0 {
		return model.Order{}, mongo.ErrNoDocuments
	}
	return model.Order{}, nil
}

func (o *OrderRepo) Delete(c context.Context, id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := o.DB.Collection("order").DeleteOne(c, bson.M{"_id": ID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
