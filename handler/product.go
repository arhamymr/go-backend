package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"networking/database"
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

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product = database.Product{
		Name: "hellow",
		Description: "desc",
		Price: 10000,
		ProductPicture: []string{"http-link"},
		Discount: 20,
		DetailProduct: "hellow",
	}

	database.DB.Create(&product);

	jsonData, err := json.Marshal(product)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []database.Product
	database.DB.Find(&products);
	fmt.Println(products, &database.DB, "- inside getPRoductd,", database.DB)
	
	jsonData, err := json.Marshal(products)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
