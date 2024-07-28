package repositories

import (
	"context"

	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	ModelDb *mongo.Collection
	*BaseRepository[models.UserModel]
}

type UserRepositoryInterface interface {
	BaseRepositoryInterface[models.UserModel]
	GetByEmail(email string) (*models.UserModel, error)
}

func NewUserRepository(modelDb *mongo.Collection) UserRepositoryInterface {
	return &UserRepository{ModelDb: modelDb, BaseRepository: NewBaseRepository[models.UserModel](modelDb)}
}


func (r *UserRepository) GetByEmail(email string) (*models.UserModel, error) {

	data := &models.UserModel{}

	err := r.ModelDb.FindOne(context.Background(), bson.D{{Key: "email", Value: email }}).Decode(data)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}