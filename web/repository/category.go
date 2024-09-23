package repository

import (
	"context"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepoInterface interface {
	FindByID(c context.Context, id string) (model.Category, error)
	Create(c context.Context, category model.Category) (model.Category, error)
	Update(c context.Context, category model.Category) (model.Category, error)
	Delete(c context.Context, id string) error
}

type CategoryRepo struct {
	DB *mongo.Database
}

func NewCategoryRepo(db *mongo.Database) CategoryRepoInterface {
	return &CategoryRepo{DB: db}
}

func (ca *CategoryRepo) FindByID(c context.Context, id string) (model.Category, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Category{}, err
	}
	var category model.Category
	err = ca.DB.Collection("category").FindOne(c, bson.M{"_id": ID}).Decode(&category)
	if err != nil {
		return model.Category{}, err
	}
	return model.Category{}, nil
}

func (ca *CategoryRepo) Create(c context.Context, category model.Category) (model.Category, error) {
	result, err := ca.DB.Collection("category").InsertOne(c, category)
	if err != nil {
		return model.Category{}, err
	}
	category.ID = result.InsertedID.(primitive.ObjectID)
	return category, nil
}

func (ca *CategoryRepo) Update(c context.Context, category model.Category) (model.Category, error) {
	result, err := ca.DB.Collection("category").UpdateOne(c, bson.M{"_id": category.ID}, bson.M{
		"$set": bson.M{
			"name": category.Name,
		},
	})
	if err != nil {
		return model.Category{}, err
	}
	if result.MatchedCount == 0 {
		return model.Category{}, mongo.ErrNoDocuments
	}
	return model.Category{}, nil
}

func (ca *CategoryRepo) Delete(c context.Context, id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := ca.DB.Collection("category").DeleteOne(c, bson.M{"_id": ID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return err
	}
	return nil
}
