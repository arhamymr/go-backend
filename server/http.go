package server

import (
	"log"
	"net/http"
	"networking/utils"
	"strconv"
)

type WebServer struct {
	routes map[string]*Route
	context *Context
	Middleware []HandlerFunc
}

func CreateWebserver() *WebServer {
	return &WebServer{
			routes: make(map[string]*Route),
		}
}

func (ws *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	ctx := NewContext(w, r)
	ws.applyMiddleware(ctx)
	if route, ok := ws.routes[r.Method+" "+r.URL.Path]; ok {
		route.applyMiddleware(ctx)
		log.Printf("Received %s request : %s", r.Method, r.URL.Path);
		route.handler(ctx)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed);
	}
}

func (ws *WebServer) addRoute(method string, url string, handler HandlerFunc) *Route {
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
	return route
}

func (ws *WebServer) AddMiddleware(middleware ...HandlerFunc) *WebServer {
	ws.Middleware = append(ws.Middleware, middleware...)
	return ws;
}

func (ws *WebServer) applyMiddleware(ctx *Context) {
	for _, m := range ws.Middleware {
		m(ctx)
	}
}

func (ws *WebServer) GET(url string, handler HandlerFunc) *Route {
	return ws.addRoute(http.MethodGet, url, handler )
}
func (ws *WebServer) PUT(url string, handler HandlerFunc) *Route {
	return ws.addRoute(http.MethodPut, url, handler )
}

func (ws *WebServer) PATCH(url string, handler HandlerFunc) *Route{
	return ws.addRoute(http.MethodPatch, url, handler )
}

func (ws *WebServer) DELETE(url string, handler HandlerFunc) *Route {
	return ws.addRoute(http.MethodDelete, url, handler )
}

func (ws *WebServer) POST(url string, handler HandlerFunc) *Route {
	return ws.addRoute(http.MethodPost, url, handler )
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
