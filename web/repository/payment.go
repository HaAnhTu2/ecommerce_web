package repository

import (
	"context"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepoInterface interface {
	FindByID(c context.Context, id string) (model.Payment, error)
	Create(c context.Context, payment model.Payment) (model.Payment, error)
	Update(c context.Context, payment model.Payment) (model.Payment, error)
	Delete(c context.Context, id string) error
}

type PaymentRepo struct {
	DB *mongo.Database
}

func NewPaymentRepo(db *mongo.Database) PaymentRepoInterface {
	return &PaymentRepo{DB: db}
}

func (pm *PaymentRepo) FindByID(c context.Context, id string) (model.Payment, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Payment{}, err
	}
	var payment model.Payment
	err = pm.DB.Collection("payment").FindOne(c, bson.M{"_id": ID}).Decode(&payment)
	if err != nil {
		return model.Payment{}, err
	}
	return model.Payment{}, nil
}

func (pm *PaymentRepo) Create(c context.Context, payment model.Payment) (model.Payment, error) {
	result, err := pm.DB.Collection("payment").InsertOne(c, payment)
	if err != nil {
		return model.Payment{}, err
	}
	payment.ID = result.InsertedID.(primitive.ObjectID)
	return payment, err
}

func (pm *PaymentRepo) Update(c context.Context, payment model.Payment) (model.Payment, error) {
	result, err := pm.DB.Collection("payment").UpdateOne(c, bson.M{"_id": payment.ID}, bson.M{
		"$set": bson.M{
			"payment_method": payment.PaymentMethod,
			"status":         payment.Status,
		},
	})
	if err != nil {
		return model.Payment{}, err
	}
	if result.MatchedCount == 0 {
		return model.Payment{}, mongo.ErrNoDocuments
	}
	return model.Payment{}, nil
}

func (pm *PaymentRepo) Delete(c context.Context, id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := pm.DB.Collection("payment").DeleteOne(c, bson.M{"_id": ID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
