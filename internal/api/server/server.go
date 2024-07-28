package server

import (
	"context"

	"github.com/joaops3/go-olist-challenge/internal/api/router"
	config "github.com/joaops3/go-olist-challenge/internal/configs"
	"github.com/joaops3/go-olist-challenge/internal/data/db"
)

type HttpServer struct {
}

func (s *HttpServer) Serve() {
	logger := config.GetLogger("MAIN")
	config.LoadEnvFile()
	dbClient, err := db.InitDb()
	defer dbClient.Disconnect(context.Background())

	if err != nil {
		logger.Errorf(err.Error())
		panic("error main")
	}
	router.Initialize()
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}