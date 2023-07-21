package handler

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

func GetPostArticle(w http.ResponseWriter, r *http.Request) {
	post := Post{
		Title: "Title here",
		Content: "content here",
	}

	jsonData, err := json.Marshal(post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}


func SubmitPostArticle(w http.ResponseWriter, h http.Request) {
	
}
