package httpServer

import (
	"fmt"
	"net/http"
)

type eventHandler func(http.ResponseWriter, *http.Request)

type Handler struct {
	http.Handler
	EventHandlers map[string]eventHandler
}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.Get(res, req)
	case "POST":
		h.Post(res, req)
	case "PUT":
		h.Put(res, req)
	case "PATCH":
		h.Patch(res, req)
	case "DELETE":
		h.Delete(res, req)
	default:
		h.Unsupported(res, req)
	}
}

func (h Handler) ToBeImplemented(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintln(res, "{\"message\": \"Coming soon\"}")
	if err != nil {
		panic(err)
	}
}

func (h Handler) Get(res http.ResponseWriter, req *http.Request) {
	h.ToBeImplemented(res, req)
}

func (h Handler) Post(res http.ResponseWriter, req *http.Request) {
	h.ToBeImplemented(res, req)
}

func (h Handler) Put(res http.ResponseWriter, req *http.Request) {
	h.ToBeImplemented(res, req)
}

func (h Handler) Patch(res http.ResponseWriter, req *http.Request) {
	h.ToBeImplemented(res, req)
}

func (h Handler) Delete(res http.ResponseWriter, req *http.Request) {
	h.ToBeImplemented(res, req)
}

func (h Handler) Unsupported(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(res, "{\"message\": \"Method Not Supported\"}")
}
