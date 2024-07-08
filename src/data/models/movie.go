package models

import (
	"github.com/joaops3/go-olist-challenge/src/data/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieModel struct {
	BaseModel `bson:",inline"`
	Name string `bson:"name,omitempty"   json:"name" csv:"title"` 
	Genre string `bson:"genre,omitempty"   json:"genre" csv:"genres"`
}

func GetDbMovieCollection() *mongo.Collection {
	Db := db.GetDb()
	model := Db.Collection("movies")
	return model
}


func NewMovieModel(name string, genre string) *MovieModel{
	v := &MovieModel{Name: name, Genre: genre}
	v.InitBaseModel()
	return v
}