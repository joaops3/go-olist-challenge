package server

import (
	"github.com/joaops3/go-olist-challenge/src/api/router"
	config "github.com/joaops3/go-olist-challenge/src/configs"
	"github.com/joaops3/go-olist-challenge/src/data/db"
)

type HttpServer struct {
}

func (s *HttpServer) Serve() {
	logger := config.GetLogger("MAIN")
	config.LoadEnvFile()
	_, err := db.InitDb()
	// defer dbClient.Disconnect(context.Background())

	if err != nil {
		logger.Errorf(err.Error())
		panic("error main")
	}
	router.Initialize()
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}