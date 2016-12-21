package httpServer

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cpg1111/pr-show-metrics-pipeline/www"
)

type StaticHandler struct {
	http.Handler
}

func (s StaticHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(res, "{\"message\": \"method unsupported\"}")
		return
	}
	pathSplit := strings.Split(req.URL.Path, "/")
	path := strings.Join(pathSplit[1:], "/")
	res.WriteHeader(http.StatusOK)
	asset, err := www.Asset(path)
	if err != nil {
		panic(err)
	}
	res.Write(asset)
}
