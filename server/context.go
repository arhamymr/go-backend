package server

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
	Middleware []HandlerFunc
}

func NewContext (w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response: w,
		Request: r,
	}
}

func (ctx *Context) JSON(data interface{}) error {
	jsonData, err := json.Marshal(data)
	ctx.Response.Header().Set("Content-Type", "application/json")
	ctx.Response.WriteHeader(http.StatusOK)
	ctx.Response.Write(jsonData)
	return err
}

