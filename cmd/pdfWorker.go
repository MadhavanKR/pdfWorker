package main

import (
	"github.com/MadhavanKR/pdfWorker/pkg/server"
)

func main() {
	httpServer := server.GetHttpServer()
	httpServer.ListenAndServe()
}
