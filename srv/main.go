package main

import (
	"log"
	"os"

	"github.com/cpg1111/pr-show-metrics-pipeline/pkg/httpServer"
)

func main() {
	log.Println("starting server")
	server := httpServer.New(os.Getenv("PROTO"), os.Getenv("HOST"), os.Getenv("PORT"))
	if len(os.Getenv("USE_TLS")) > 0 {
		server.ListenAndServeTLS(os.Getenv("CERT_PATH"), os.Getenv("KEY_PATH"))
	} else {
		server.ListenAndServe()
	}
}
