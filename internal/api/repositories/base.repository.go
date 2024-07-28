package repositories

import (
	"context"
	"errors"
	"math"
	"reflect"
	"time"

	"github.com/joaops3/go-olist-challenge/internal/api/dtos"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseRepositoryInterface[T models.Identifiable] interface {
	BaseGetPaginated(query *dtos.PaginationDto) ([]*T, error)
	BaseGetById(id string) (*T, error)
	BaseSave(v *T) (*T, error)
	BaseCreate(v *T) (*T, error)
	BaseUpdate(id string , dto any) (bool, error)
	BaseDelete(id string) error
}

type BaseRepository[T models.Identifiable] struct {
	ModelDb *mongo.Collection
}

func NewBaseRepository[T models.Identifiable](db  *mongo.Collection) *BaseRepository[T]{
	return &BaseRepository[T]{
			ModelDb: db,
	}
}

func (b *BaseRepository[T]) BaseGetPaginated(query *dtos.PaginationDto) ([]*T, error) {
	
	where := bson.D{bson.E{Key: "deleted_at", Value: nil}}
	options := options.Find()
	options.SetLimit(query.PageSize)


	total, err := b.ModelDb.CountDocuments(context.Background(), where) 

	if err != nil {
		return nil, err
	}
	pages := max(math.Ceil(float64((total / query.PageSize))), 1)
	if query.Current > int64((pages)) {
		query.Current = int64(pages)
	}

	skip := query.PageSize * query.Current - query.Current
	options.SetSkip(skip)

	cursor, err := b.ModelDb.Find(context.Background(), where, options)

	if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())
	
	var results []*T
    if err = cursor.All(context.Background(), &results); err != nil {
        return nil, err
    }
	return results, nil
}

func (b *BaseRepository[T]) BaseGetById(id string) (*T, error) {
    var result T
    _id, err := primitive.ObjectIDFromHex(id)

	
    if err != nil {
        return nil, err
    }
    err = b.ModelDb.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&result)

    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, err
    }
    return &result, nil
}


func (b *BaseRepository[T]) BaseSave(v *T) (*T, error) {
   
	model := *v
	
    filter := bson.M{"_id": model.GetID()}
	model.SetUpdatedAt(time.Now())
    update := bson.M{"$set": model }
	
	opts := options.Update().SetUpsert(true)
    result, err := b.ModelDb.UpdateOne(context.TODO(), filter, update, opts)
    if err != nil {
        return nil, err
    }
	if result.MatchedCount == 0 && result.UpsertedCount == 0 {
		return nil, errors.New("no matched document found for update")
	}
    return v, nil
}

func (b *BaseRepository[T]) BaseCreate(v *T) (*T, error) {
	
	_, err := b.ModelDb.InsertOne(context.TODO(), v)
	if err != nil {
		return nil, err
	}
	return v, nil
}


func (b *BaseRepository[T]) BaseUpdate(id string,  dto any) (bool, error) {
	
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return false, err
	}

	
	filter := bson.D{{Key: "_id", Value: objectID}}

	
	updatedFields := bson.D{}

	dtoValue := reflect.ValueOf(dto)
	dtoType := dtoValue.Type()

	for i := 0; i < dtoValue.NumField(); i++ {
		field := dtoValue.Field(i)
		fieldType := dtoType.Field(i)
		fieldName := fieldType.Name
		if field.CanInterface() {
			updatedFields = append(updatedFields, bson.E{Key: fieldName, Value: field.Interface()})
		}
	}

	
	updatedFields = append(updatedFields, bson.E{Key: "updated_at", Value: time.Now()})

	
	update := bson.M{"$set": updatedFields}

	
	result, err := b.ModelDb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return false, err
	}

	if result.MatchedCount == 0 {
		return false, errors.New("no matched document found for update")
	}

	return true, nil
}


func (b *BaseRepository[T]) BaseDelete(id string) error {
	
	_id, erro :=  primitive.ObjectIDFromHex(id)
	if erro != nil {
		return nil
	}
	
	where := bson.D{{Key: "_id", Value:_id}}

	updatedFields := bson.D{}
	updatedFields = append(updatedFields, bson.E{Key: "deleted_at", Value: time.Now()})

	command := bson.M{"$set": updatedFields}

	_, err :=  b.ModelDb.UpdateOne(context.Background(), where, command)
	if err != nil {
		return  err
	}
	return  nil
}