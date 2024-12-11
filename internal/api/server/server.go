package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/joaops3/go-olist-challenge/internal/api/router"
	config "github.com/joaops3/go-olist-challenge/internal/configs"
	"github.com/joaops3/go-olist-challenge/internal/data/db"
)

type HttpServer struct {
}

func (s *HttpServer) Serve() ( *gin.Engine, error){
	logger := config.GetLogger("MAIN")
	config.LoadEnvFile()
	dbClient, err := db.InitDb() 
	defer dbClient.Disconnect(context.Background())

	if err != nil {
		logger.Errorf(err.Error())
		return nil, err
	}
	engine, err := router.Initialize()
	if err != nil {
		logger.Errorf(err.Error())
		return nil, err
	}

	return engine, nil
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}