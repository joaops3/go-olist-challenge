package main

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joaops3/go-olist-challenge/internal/api/server"
)

func main() {
	s := server.NewHttpServer()
	
	go func() {
		if _, err := s.Serve(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err) 
			}
		}
	}()

	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	// fmt.Println("Shutting down server...")

	// if err := s.Shutdown(); err != nil {
	// 	fmt.Printf("Error during server shutdown: %v\n", err)
	// } else {
	// 	fmt.Println("Server gracefully stopped.")
	// }
}