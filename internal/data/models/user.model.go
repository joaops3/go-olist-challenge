package models

import (
	"github.com/joaops3/go-olist-challenge/internal/data/db"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_USER string = "users"

type UserModel struct {
	*BaseModel `bson:",inline"`
	Email      string `json:"email" bson:"email,omitempty"  `
	Password   string `json:"password" bson:"password,omitempty"`
	ProfileImg  *S3FileModel `json:"profileImg" bson:"profileImg,omitempty"`
}

type JwtResponse struct {
	Id   string `json:"_id"`
	Token string `json:"token"`
}

func GetDbUserCollection() *mongo.Collection {
	Db := db.GetDb()
	model := Db.Collection(COLLECTION_MOVIE)
	return model
}

func NewUserModel(email string, password string) *UserModel {
	v := &UserModel{Email: email, Password: password, BaseModel: &BaseModel{}}
	v.InitBaseModel()
	return v
}