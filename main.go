package main

import (
	"networking/database"
	"networking/handler"
	"networking/server"
)


func main() {
	database.InitDatabase();
	r := server.CreateWebserver()
	r.GET("/", handler.GetPostArticle)
	r.GET("/users", handler.GetPostArticle)
	r.GET("/products", handler.GetProducts)
	r.POST("/products", handler.AddProduct)

	const PORT = 8080;
	server.StartServer(PORT, r)	
}


