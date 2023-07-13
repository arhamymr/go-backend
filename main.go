package main

import (
	"networking/handler"
	"networking/server"
)

func main() {

	r := server.CreateWebserver()
	r.GET("/", handler.GetPost)
	r.GET("/users", handler.GetPost)

	const PORT = 8080;
	server.StartServer(PORT, r)	
}


