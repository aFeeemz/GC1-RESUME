package repository

import (
	"GC1/models"
	"GC1/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FieldRepo interface {
	GetAllFields() ([]models.Field, error)
}

func (r *Repo) GetAllFields() ([]models.Field, error) {
	var fields []models.Field
	cursor, err := r.DB.Collection("fields").Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, utils.ErrQuerying
	}

	for cursor.Next(context.TODO()) {
		var field models.Field
		if err := cursor.Decode(&field); err != nil {
			return nil, utils.ErrQuerying
		}
		fields = append(fields, field)
	}
	return fields, nil
}

func (r *Repo) CheckField(field_id primitive.ObjectID) (bool, error) {
	var result bson.M

	err := r.DB.Collection("fields").FindOne(context.TODO(), bson.M{"_id": field_id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User not found
			return false, nil
		}
		return false, utils.ErrQuerying
	}

	return true, err
}
