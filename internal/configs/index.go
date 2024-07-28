package config

import (
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var	( 
	client *mongo.Client 
	logger *Logger
)




func GetLogger(p string) *Logger {
	// Initialize Logger
	logger = NewLogger(p)
	return logger
}

func LoadEnvFile(){
	err := godotenv.Load()
	
	if err != nil {
	  logger.Error("Error loading .env file")
	  panic(err.Error())
	}
  
}