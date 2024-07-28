package models

import (
	"github.com/joaops3/go-olist-challenge/internal/data/db"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_MOVIE string = "movies"

type MovieModel struct {
	*BaseModel `bson:",inline"`
	Name string `bson:"name,omitempty"   json:"name" csv:"title"` 
	Genre string `bson:"genre,omitempty"   json:"genre" csv:"genres"`
}

func GetDbMovieCollection() *mongo.Collection {
	Db := db.GetDb()
	model := Db.Collection(COLLECTION_MOVIE)
	return model 
}


func NewMovieModel(name string, genre string) *MovieModel{
	v := &MovieModel{Name: name, Genre: genre, BaseModel: &BaseModel{},}
	v.InitBaseModel()
	return v
}




