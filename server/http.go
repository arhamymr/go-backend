package server

import (
	"log"
	"net/http"
	"networking/utils"
	"strconv"
)

type (
	Route struct {
		url string
		method string
		handler HandlerFunc
		params map[string][]string
		body []byte
	}	

	
	WebServer struct {
		// mux *http.ServeMux
		routes map[string]*Route
		context *Context
	}

	HandlerFunc func (ctx *Context)
)



func CreateWebserver() *WebServer {
	return &WebServer{
			routes: make(map[string]*Route),
		}
}


func (ws *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if route, ok := ws.routes[r.Method+" "+r.URL.Path]; ok {
		log.Printf("Received GET request : %s %s", r.Method, r.URL.Path);
		ctx := NewContext(w, r)
		route.handler(ctx)
		
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed);
	}
}

func (ws *WebServer) addRoute(method string, url string, handler HandlerFunc) {
	if ws.routes == nil {
		ws.routes = make(map[string]*Route)
	} 
	route := &Route{
		url: url,
		method: method,
		handler: handler,
	}
	key := method + " " + url
	ws.routes[key] = route
}

func (ws *WebServer) GET(url string, handler HandlerFunc){
	ws.addRoute(http.MethodGet, url, handler )
}
func (ws *WebServer) PUT(url string, handler HandlerFunc){
	ws.addRoute(http.MethodPut, url, handler )
}

func (ws *WebServer) PATCH(url string, handler HandlerFunc){
	ws.addRoute(http.MethodPatch, url, handler )
}

func (ws *WebServer) DELETE(url string, handler HandlerFunc){
	ws.addRoute(http.MethodDelete, url, handler )
}

func (ws *WebServer) POST(url string, handler HandlerFunc){
	ws.addRoute(http.MethodPost, url, handler )
}


func StartServer(port uint16, handler http.Handler) {
	// check valid port number 
	err := utils.CheckValidPortNumber(port)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server is running on port:", port)
	err = http.ListenAndServe(":"+strconv.Itoa(int(port)), handler)

	if err != nil {
		log.Fatal("Failed to start web server: %v", err)
	}
}
