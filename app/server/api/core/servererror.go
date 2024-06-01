package core

import (
	"fmt"
	"log"
	"net/http"
)

type ServerErrorHandler struct {
}

func (h ServerErrorHandler) WithError(err error, w http.ResponseWriter, r *http.Request) {
	log.Println(err)
	h.ServeHTTP(w, r)
}

func (h ServerErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Server error 500")
}

var ServerError ServerErrorHandler
