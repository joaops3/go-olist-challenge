package repositories

import (
	"context"
	"time"

	"github.com/joaops3/go-olist-challenge/src/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MovieRepository struct {
	ModelDb *mongo.Collection
}

type MovieRepositoryInterface interface{
	BaseRepositoryInterface[models.MovieModel]
}

func NewMovieRepository(modelDb *mongo.Collection) *MovieRepository {
	return &MovieRepository{ModelDb: modelDb}
}


func (r *MovieRepository) Save(model *models.MovieModel) (*models.MovieModel, error) {
	filter := bson.M{"_id": model.ID}
	update := bson.M{"$set": model }
	opts := options.FindOneAndUpdate().SetUpsert(true)

	result := r.ModelDb.FindOneAndUpdate(context.Background(), filter, update, opts)

	if result.Err() != nil {
		return nil, result.Err()
	}
	return model,nil
}

func (r *MovieRepository) GetById(id string)(*models.MovieModel, error) {
	customer := &models.MovieModel{}

	_id, erro :=  primitive.ObjectIDFromHex(id)

	if erro != nil {
		return nil, erro
	}
	
	where := bson.D{bson.E{Key: "_id", Value: _id}}
	err := r.ModelDb.FindOne(context.Background(), where).Decode(customer)

	if err != nil {
		return nil, err
	}
	return customer, nil
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

func (r *MovieRepository) GetPaginated() ([]*models.MovieModel, error) {

	movies := []*models.MovieModel{}
    where := bson.D{}
	options := options.Find()
	options.SetLimit(5)
	cursor, err := r.ModelDb.Find(context.Background(), where, options)

	if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    if err = cursor.All(context.Background(), &movies); err != nil {
        return nil, err
    }

	return movies, nil
}
