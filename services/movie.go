package service

import (
	"context"
	"go-fiber/database"
	model "go-fiber/database/models"
	"go-fiber/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveMovie(movie *model.Movie) error {
	movieModel := database.Mongo.DB.Collection("Movie")
	movie.ID = helper.GetHash(movie.Name + movie.Source)
	_, err := movieModel.InsertOne(context.Background(), movie)
	return err
}

func GetMovie(movieId string) (model.Movie, error) {
	movieModel := database.Mongo.DB.Collection("Movie")
	opts := options.FindOne()
	var movie model.Movie
	err := movieModel.FindOne(context.TODO(), bson.D{{Key: "id", Value: movieId}}, opts).Decode(&movie)
	return movie, err
}

func DeleteMovie(movieId string) error {
	movieModel := database.Mongo.DB.Collection("Movie")
	opts := options.Delete()
	_, err := movieModel.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: movieId}}, opts)
	return err
}

func GetMoviesFromGenre(genre string) ([]model.Movie, error) {
	movieModel := database.Mongo.DB.Collection("Movie")
	opts := options.Find()
	var movieList []model.Movie
	cursor, err := movieModel.Find(context.TODO(), bson.M{"genre": genre}, opts)
	cursor.All(context.TODO(), &movieList)
	return movieList, err
}

func GetMoviesFromCategory(category string) ([]model.Movie, error) {
	movieModel := database.Mongo.DB.Collection("Movie")
	opts := options.Find()
	var movieList []model.Movie
	cursor, err := movieModel.Find(context.TODO(), bson.M{"category": category}, opts)
	cursor.All(context.TODO(), &movieList)
	return movieList, err
}

func AddWatchlist(watchlistBody *model.Watchlist) error {
	watchlistModel := database.Mongo.DB.Collection("Watchlist")
	_, err := watchlistModel.InsertOne(context.Background(), watchlistBody)
	return err
}
