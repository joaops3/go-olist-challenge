package main

import "github.com/joaops3/go-olist-challenge/src/api/server"

func main() { 
	
	s := server.NewHttpServer() 
	s.Serve()
}