package repository

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoInterface interface {
	FindByID(ctx context.Context, id string) (model.User, error)
	FindByEmail(c context.Context, email string) (model.User, error)
	Create(c context.Context, user model.User) (model.User, error)
	Update(c context.Context, user model.User) (model.User, error)
	Delete(ctx context.Context, id string) error
	SaveToken(user model.User) (string, error)
}

type UserRepo struct {
	db *mongo.Database
}

func NewUserRepo(db *mongo.Database) UserRepoInterface {
	return &UserRepo{db: db}
}

func (u *UserRepo) FindByID(ctx context.Context, id string) (model.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.User{}, errors.New("invalid user ID")
	}

	var user model.User
	err = u.db.Collection("users").FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, err
	}
	return user, nil
}
func (u *UserRepo) FindByEmail(c context.Context, email string) (model.User, error) {
	var user model.User
	err := u.db.Collection("users").FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.User{}, err
		}
		return model.User{}, err
	}
	return user, nil
}

func (u *UserRepo) Create(c context.Context, user model.User) (model.User, error) {
	result, err := u.db.Collection("users").InsertOne(c, user)
	if err != nil {
		return model.User{}, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (u *UserRepo) Update(c context.Context, user model.User) (model.User, error) {
	result, err := u.db.Collection("users").UpdateOne(c, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			"username": user.UserName,
			"password": user.Password,
			"email":    user.Email,
			"address":  user.Address,
			"phone":    user.Phone,
		},
	})
	if err != nil {
		return model.User{}, err
	}
	if result.MatchedCount == 0 {
		return model.User{}, mongo.ErrNoDocuments
	}
	return user, nil
}

func (u *UserRepo) Delete(ctx context.Context, id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := u.db.Collection("users").DeleteOne(ctx, bson.M{"_id": ID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return err
	}
	return nil
}

func (u *UserRepo) SaveToken(user model.User) (string, error) {
	secret := os.Getenv("SECRET_KEY")
	expired_At := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"sub":  user.Email,
		"role": user.Role,
		"exp":  expired_At.Unix(),
	})
	return token.SignedString([]byte(secret))
}
