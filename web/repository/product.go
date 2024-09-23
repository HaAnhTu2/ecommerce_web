package repository

import (
	"context"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepoInterface interface {
	FindByID(c context.Context, id string) (model.Product, error)
	GetAll(ctx context.Context) ([]model.Product, error)
	Create(c context.Context, product model.Product) (model.Product, error)
	Update(c context.Context, product model.Product) (model.Product, error)
	Delete(c context.Context, id string) error
}

type ProductRepo struct {
	DB *mongo.Database
}

func NewProductRepo(db *mongo.Database) ProductRepoInterface {
	return &ProductRepo{DB: db}
}

func (p *ProductRepo) FindByID(c context.Context, id string) (model.Product, error) {
	var product model.Product
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Product{}, err
	}
	err = p.DB.Collection("products").FindOne(c, bson.M{"_id": ID}).Decode(&product)
	if err != nil {
		return model.Product{}, err
	}
	return model.Product{}, nil
}
func (p *ProductRepo) GetAll(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	result, err := p.DB.Collection("products").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := result.All(context.Background(), &products); err != nil {
		return nil, err
	}
	for _, item := range products {
		products = append(products, model.Product{
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Stock:       item.Stock,
			Category:    item.Category,
		})
	}
	return products, nil
}

func (p *ProductRepo) Create(c context.Context, product model.Product) (model.Product, error) {
	result, err := p.DB.Collection("products").InsertOne(c, product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = result.InsertedID.(primitive.ObjectID)
	return product, nil
}

func (p *ProductRepo) Update(c context.Context, product model.Product) (model.Product, error) {
	result, err := p.DB.Collection("products").UpdateOne(c, bson.M{"_id": product.ID}, bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"description": product.Description,
			"stock":       product.Stock,
			"price":       product.Price,
			"category":    product.Category,
		},
	})
	if err != nil {
		return model.Product{}, err
	}
	if result.MatchedCount == 0 {
		return model.Product{}, err
	}
	return model.Product{}, nil
}

func (p *ProductRepo) Delete(c context.Context, id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := p.DB.Collection("products").DeleteOne(c, bson.M{"_id": ID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return err
	}
	return nil
}
