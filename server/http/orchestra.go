package http

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	address string
	mux     *http.ServeMux
}

func NewServer(ip, port string) *Server {
	return &Server{
		address: fmt.Sprintf("%s:%s", ip, port),
		mux:     http.NewServeMux(),
	}
}

func (s *Server) LoadEndpoints() {
	s.mux.Handle("", &ExpressionsApi{})
}

func (s *Server) StartServer(ip, port string) {

	log.Fatal(http.ListenAndServe(s.address, s.mux))
}
