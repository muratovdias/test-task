package repository

import (
	"context"
	"fmt"
	"log"
	"user/internal/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	CreateUser(models.User) (models.User, error)
	GetUser(email string) (models.User, error)
	FindUser(email string) (int, error)
}

type UserRepo struct {
	collection mongo.Collection
}

func NewUserRepo(collection *mongo.Collection) *UserRepo {
	return &UserRepo{
		collection: *collection,
	}
}

func (r *UserRepo) CreateUser(user models.User) (models.User, error) {
	_, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return models.User{}, err
	}
	log.Println("User created successfully")
	return user, nil
}

func (r *UserRepo) FindUser(email string) (int, error) {
	filter := bson.M{"email": email}
	count, err := r.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	if count >= 1 {
		return int(count), nil
	}
	return 0, nil
}

func (r *UserRepo) GetUser(email string) (models.User, error) {
	filter := bson.M{"email": email}
	var user models.User
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("repository: %w", err)
	}
	return user, nil
}
