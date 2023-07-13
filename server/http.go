package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type WebServer struct {
	mux *http.ServeMux
}

func (ws *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	ws.mux.ServeHTTP(w,r)
}

func (ws *WebServer) GET(url string, handler http.HandlerFunc){
	ws.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return 
		}
		handler(w, r)
	})
}

func CreateWebserver() *WebServer {
	return &WebServer{
			mux: http.NewServeMux(),
		}
}

func checkValidPortNumber(port uint16) error {
	const MAX_PORT_RANGE = 65535
	if port > MAX_PORT_RANGE {
		return errors.New("invalid port number")
	}
	return nil
}

func handlePOSTRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handling POST request...")
	// Custom logic for handling POST request
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
