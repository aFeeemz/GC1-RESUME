package repository

import (
	"GC1/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionRepo interface {
	CreateTransaction(t *models.Transaction) (interface{}, error)
}

type Repo struct {
	DB *mongo.Database
}

func (r *Repo) CreateTransaction(t *models.Transaction) (interface{}, error) {

	// Check if user exists
	isUserExists, err := r.CheckUser(t.UserID)
	if err != nil || !isUserExists {
		return nil, err
	}

	// Insert Transaction
	result, err := r.DB.Collection("transactions").InsertOne(context.TODO(), t)
	if err != nil {
		return nil, err
	}

	return result, err

}
