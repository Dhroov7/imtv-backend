package service

import (
	"context"
	"go-fiber/database"
	model "go-fiber/database/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Save(user *model.User) error {
	var userModel = database.Mongo.DB.Collection("user")
	_, err := userModel.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(email string) (model.User, error) {
	var userModel = database.Mongo.DB.Collection("user")
	opts := options.FindOne()
	var user model.User
	err := userModel.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}, opts).Decode(&user)
	return user, err
}
