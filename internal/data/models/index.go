package models

import (
	"github.com/joaops3/go-olist-challenge/internal/data/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var Db *mongo.Database


func InitDb() {
	Db = db.GetDb()
}

