package handler

import (
	"fmt"
	"net/http"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hellow World");
}

