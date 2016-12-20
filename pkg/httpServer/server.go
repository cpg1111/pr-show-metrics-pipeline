package httpServer

import (
	"fmt"
	"net/http"
)

type Server struct {
	http.Server
}

func New(proto, host, port string) *Server {
	mux := http.NewServeMux()
	mux.Handle("/assets/", StaticHandler{})
	mux.Handle("/events/", Handler{})
	return &Server{
		http.Server{
			Addr:    fmt.Sprintf("%s://%s:%s", proto, host, port),
			Handler: mux,
		},
	}
}
