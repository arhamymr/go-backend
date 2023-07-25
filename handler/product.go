package handler

import (
	"fmt"
	"net/http"
	"networking/database"
	"networking/server"
)

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ProductPicture []string `json:"product_picture"`
	Discount int `json:"discount"`
	DetailProduct string `json:"detail_product"`
	Likes int `json:"likes"`
}

func AddProduct(ctx *server.Context) {
	var product = database.Product{
		Name: "hellow",
		Description: "desc",
		Price: 10000,
		ProductPicture: []string{"http-link"},
		Discount: 20,
		DetailProduct: "hellow",
	}

	database.DB.Create(&product);
	err := ctx.JSON(product)
	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return 
	}
}


func GetProducts(ctx *server.Context) {
	var products []database.Product
	database.DB.Find(&products);
	fmt.Println(products, &database.DB, "- inside getPRoductd,", database.DB)
	
	err := ctx.JSON(products)

	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return 
	}
}
