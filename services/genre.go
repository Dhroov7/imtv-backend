package service

import (
	"context"
	"go-fiber/database"
	model "go-fiber/database/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveGenre(genreData *model.Genre) error {
	genreModel := database.Mongo.DB.Collection("Genre")
	_, err := genreModel.InsertOne(context.Background(), genreData)
	return err
}

func GetGenres() ([]model.Genre, error) {
	genreModel := database.Mongo.DB.Collection("Genre")
	opts := options.Find()
	var genreList []model.Genre
	cursor, err := genreModel.Find(context.TODO(), bson.M{}, opts)
	cursor.All(context.TODO(), &genreList)
	return genreList, err
}
