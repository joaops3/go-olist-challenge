package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joaops3/go-olist-challenge/internal/api/router"
	config "github.com/joaops3/go-olist-challenge/internal/configs"
	"github.com/joaops3/go-olist-challenge/internal/data/db"
)

type HttpServer struct {
	httpServer *http.Server
}

func (s *HttpServer) Serve() ( *gin.Engine, error){
	logger := config.GetLogger("MAIN")
	config.LoadEnvFile()
	dbClient, err := db.InitDb() 
	if err != nil {
		logger.Errorf(err.Error())
		return nil, err
	}
	defer dbClient.Disconnect(context.Background())
	engine, err := router.Initialize()
	if err != nil {
		logger.Errorf(err.Error())
		return nil, err
	}

	s.httpServer = &http.Server{
		Addr:    ":3333",
		Handler: engine,
	}

	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Errorf("HTTP server error: %v", err)
		return nil, err
	}

	return engine, nil
}

func (s *HttpServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}