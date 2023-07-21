package server

import (
	"errors"
	"log"
	"net/http"
	"strconv"
)

type WebServer struct {
	// mux *http.ServeMux
	routes map[string]http.HandlerFunc
}

func (ws *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if handler, ok := ws.routes[r.Method+" "+r.URL.Path]; ok {
		log.Printf("Received GET request : %s %s", r.Method, r.URL.Path);
		handler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed);
	}
}

func (ws *WebServer) addRoute(method string, url string, handler http.HandlerFunc) {
	if ws.routes == nil {
		ws.routes = make(map[string]http.HandlerFunc)
	} 
	key := method + " " + url
	ws.routes[key] = handler
}

func (ws *WebServer) GET(url string, handler http.HandlerFunc){
	ws.addRoute(http.MethodGet, url, handler )
}
func (ws *WebServer) PUT(url string, handler http.HandlerFunc){
	ws.addRoute(http.MethodPut, url, handler )
}

func (ws *WebServer) PATCH(url string, handler http.HandlerFunc){
	ws.addRoute(http.MethodPatch, url, handler )
}

func (ws *WebServer) DELETE(url string, handler http.HandlerFunc){
	ws.addRoute(http.MethodDelete, url, handler )
}

func (ws *WebServer) POST(url string, handler http.HandlerFunc){
	ws.addRoute(http.MethodPost, url, handler )
}


func CreateWebserver() *WebServer {
	return &WebServer{
			routes: make(map[string]http.HandlerFunc),
		}
}


func checkValidPortNumber(port uint16) error {
	const MAX_PORT_RANGE = 65535
	if port > MAX_PORT_RANGE {
		return errors.New("invalid port number")
	}
	return nil
}

func StartServer(port uint16, handler http.Handler) {
	// check valid port number 
	err := checkValidPortNumber(port)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server is running on port:", port)
	err = http.ListenAndServe(":"+strconv.Itoa(int(port)), handler)

	if err != nil {
		log.Fatal("Failed to start web server: %v", err)
	}
}
