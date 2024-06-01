package core

import (
	"fmt"
	"net/http"
)

type NotFoundHandler struct {
}

func (h NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not found")
	return
}

var NotFound NotFoundHandler
