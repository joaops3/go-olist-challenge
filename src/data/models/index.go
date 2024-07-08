package models

import (
	"github.com/joaops3/go-olist-challenge/src/data/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var Db *mongo.Database


func InitDb() {
	Db = db.GetDb()
}

