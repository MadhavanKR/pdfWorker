package server

import (
	"github.com/MadhavanKR/pdfWorker/pkg/apisurface"
	"github.com/gorilla/mux"
	"net/http"
)

func GetHttpServer() *http.Server{
	httpServer := &http.Server{Addr: ":2527"}
	httpServer.Handler = getServerMux()
	return httpServer
}

func getServerMux() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloWorld).Methods("GET")
	router.HandleFunc("/uploadFiles", apisurface.ConvertImagesToPdfHandler).Methods( "POST")
	return router
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}