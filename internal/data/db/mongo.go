package db

import (
	"context"
	"os"
	"time"

	config "github.com/joaops3/go-olist-challenge/internal/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeMongo() (*mongo.Client, error) {
	logger := config.NewLogger("MONGO")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	MONGO_URL := os.Getenv("MONGO_URL")
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Errorf("ERROR MONGO: %v", err)
		client.Disconnect(ctx)
		panic(err)
	}
	
	if err != nil {
		logger.Errorf("ERROR MONGO: %v", err)
		client.Disconnect(ctx)
		panic(err)
	}
	
	
	return client, nil
}

