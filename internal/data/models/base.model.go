package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	ID primitive.ObjectID `bson:"_id,omitempty"  json:"_id"`
	CreatedAt time.Time `bson:"created_at"  json:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at"  json:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at"  json:"deleted_at"`
}

type Identifiable interface {
	GetID() primitive.ObjectID
	SetUpdatedAt(t time.Time)
}

func(b *BaseModel) InitBaseModel(){
	b.ID = primitive.NewObjectID()
	b.CreatedAt = time.Now()
}

func(b *BaseModel) GetID() primitive.ObjectID{
	return b.ID
}

func(b *BaseModel) SetUpdatedAt(t time.Time){
	b.UpdatedAt = &t
}






