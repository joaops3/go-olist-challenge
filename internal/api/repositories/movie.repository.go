package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/joaops3/go-olist-challenge/internal/api/dtos"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieRepository struct {
	ModelDb *mongo.Collection
	*BaseRepository[models.MovieModel]
}

type MovieRepositoryInterface interface {
	BaseRepositoryInterface[models.MovieModel]
	GetPaginated(query *dtos.PaginationDto) ([]*models.MovieModel, error)
	GetById(id string) (*models.MovieModel, error)
	Create(v *models.MovieModel) (*models.MovieModel, error)
	Delete(id string) error
}



func NewMovieRepository(modelDb *mongo.Collection) MovieRepositoryInterface {
	return &MovieRepository{ModelDb: modelDb, BaseRepository: NewBaseRepository[models.MovieModel](modelDb)}
}

func (r *MovieRepository) GetById(id string) (*models.MovieModel, error) {
	if r.ModelDb == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	result := &models.MovieModel{}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = r.ModelDb.FindOne(context.TODO(), bson.D{bson.E{Key: "_id", Value: _id}}).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (r *MovieRepository) Create(value *models.MovieModel )(*models.MovieModel, error) {
	
	_, err := r.ModelDb.InsertOne(context.Background(), value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (r *MovieRepository) Delete(id string )( error) {
	
	_id, erro :=  primitive.ObjectIDFromHex(id)
	if erro != nil {
		return nil
	}
	
	where := bson.D{{Key: "_id", Value:_id}}

	updatedFields := bson.D{}
	updatedFields = append(updatedFields, bson.E{Key: "deleted_at", Value: time.Now()})

	command := bson.M{"$set": updatedFields}

	_, err := r.ModelDb.UpdateOne(context.Background(), where, command)
	if err != nil {
		return  err
	}
	return  nil
}

func (r *MovieRepository) GetPaginated(query *dtos.PaginationDto) ([]*models.MovieModel, error) {

	res, err := r.BaseRepository.BaseGetPaginated(query)

	if err != nil {
		return nil, err
	}

	return res, nil
}


