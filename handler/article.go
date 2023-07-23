package handler

import (
	"net/http"
	"networking/server"
)

type Post struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

func GetPostArticle(ctx *server.Context) {
	post := Post{
		Title: "Title here",
		Content: "content here",
	}

	err := ctx.JSON(post)

	if err != nil {
		http.Error(ctx.Response, err.Error(), http.StatusInternalServerError)
		return 
	}	
}


