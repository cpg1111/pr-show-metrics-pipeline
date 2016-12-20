package httpServer

import (
	"net/http"
)

type eventHandler func(ResponseWriter, *Request)

type Handler struct {
	http.Handler
	EventHandlers map[string]eventHandler
}

func (h Handler) ServeHTTP(res ResponseWriter, req *Request) {
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

func (h Handler) ToBeImplemented(res ResponseWriter, req *Request) {
	_, err := fmt.Fprintln(res, "{\"message\": \"Coming soon\"}")
	if err != nil {
		panic(err)
	}
}

func (h Handler) Get(res ResponseWriter, req *Request) {
	h.TobeImplemented(res, req)
}

func (h Handler) Post(res ResponseWriter, req *Request) {
	h.ToBeImplemented(res, req)
}

func (h Handler) Put(res ResponseWriter, req *Request) {
	h.ToBeImplemented(res, req)
}

func (h Handler) Delete(res ResponseWriter, req *Request) {
	h.ToBeImplemented(res, req)
}

func (h Handler) Unsupported(res ResponseWriter, req *Request) {
	res.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(res, "{\"message\": \"Method Not Supported\"}")
}
