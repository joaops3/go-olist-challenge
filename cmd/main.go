package main

import (
	"errors"
	"net/http"
	"os"
	"os/signal"

	"github.com/joaops3/go-olist-challenge/internal/api/server"
)

func main() { 
	
	s := server.NewHttpServer() 

	go func() {
		if err := s.Serve(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}