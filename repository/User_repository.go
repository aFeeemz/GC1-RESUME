package repository

import (
	"GC1/models"
	"GC1/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	CreateUser(t *models.User) (interface{}, error)
}

func (r *Repo) CheckUser(user_id primitive.ObjectID) (bool, error) {
	var result bson.M

	err := r.DB.Collection("users").FindOne(context.TODO(), bson.M{"_id": user_id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User not found
			return false, nil
		}
		return false, err
	}

	return true, err
}

func (r *Repo) CheckUserName(name string) (bool, error) {
	var result bson.M

	err := r.DB.Collection("users").FindOne(context.TODO(), bson.M{"name": name}).Decode(&result)
	if err != nil {
		// User not found
		if err == mongo.ErrNoDocuments {

			return false, nil
		}
		// An error occurred
		return false, utils.ErrQuerying
	}
	// User found
	return true, nil
}

func (r *Repo) CreateUser(t *models.User) (interface{}, error) {
	// Check if user exists
	isUserExists, err := r.CheckUserName(t.Name)
	if err != nil {
		return nil, utils.ErrQuerying
	}
	if isUserExists {
		return nil, utils.ErrDuplicateUser
	}

	// Insert User
	result, err := r.DB.Collection("users").InsertOne(context.TODO(), t)
	if err != nil {
		return nil, utils.ErrQuerying
	}

	return result, nil
}
