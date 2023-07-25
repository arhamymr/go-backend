package main

import (
	"fmt"
	"log"
	"networking/database"
	"networking/handler"
	"networking/server"
)


func authMiddleware(next server.HandlerFunc) server.HandlerFunc {
		// Implement authentication logic here
		// For this example, let's assume all requests pass authentication
		return func (ctx *server.Context) {
			log.Println("Authentication middleware:", ctx.Request.Method, ctx.Request.URL.Path)
			next(ctx)
		}
}

func anotherMiddleware(ctx *server.Context) {
	fmt.Println("middleware 1")
}

func anotherMiddleware2(ctx *server.Context) {
	fmt.Println("test")
}


func main() {
	database.InitDatabase();
	r := server.CreateWebserver()
	r.GET("/", handler.GetPostArticle).AddMiddleware(anotherMiddleware, anotherMiddleware2)
	r.GET("/users", handler.GetPostArticle)
	r.GET("/products", handler.GetProducts)
	r.POST("/products", handler.AddProduct)

	const PORT = 8080;
	server.StartServer(PORT, r)	
}


