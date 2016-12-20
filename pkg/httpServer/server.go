package httpServer

import (
	"fmt"
	"net/http"
)

type Server struct {
	http.Server
}

func New(proto, host string, port int) (*Server, error) {
	mux := http.NewServeMux()
	return &Server{
		Addr:    fmt.Sprintf("%s://%s:%d", proto, host, port),
		Handler: mux,
	}
}
