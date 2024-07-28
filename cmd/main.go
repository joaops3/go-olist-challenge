package main

import "github.com/joaops3/go-olist-challenge/internal/api/server"

func main() { 
	
	s := server.NewHttpServer() 
	s.Serve()
}