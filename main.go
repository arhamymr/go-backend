package main

import (
	"networking/handler"
	"networking/server"
)

func main() {
	// protocols.DialTCP("hellow from client", "0.0.0.0", "8080")
	// utils.ParseIP("127.0.0.1")
	// gui.Window(800,600)
	const PORT = 9000;
	r := server.InitWebServer(PORT)
	r.GET("/", handler.GetPost)
}

