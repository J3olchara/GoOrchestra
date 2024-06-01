package api

import (
	"fmt"
	"github.com/J3olchara/GoOrchestra/app/server/api/core"
	"github.com/J3olchara/GoOrchestra/app/server/api/expressions"
	"github.com/J3olchara/GoOrchestra/app/server/api/task"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"sync"
)

type Server struct {
	address string
	router  *mux.Router
}

var once sync.Once
var SVR *Server

func NewServer(ip, port string) *Server {
	serv := &Server{}
	created := false

	once.Do(func() {
		serv = &Server{
			address: fmt.Sprintf("%s:%s", ip, port),
			router:  mux.NewRouter(),
		}
		created = true
	})
	if !created {
		return nil
	}
	return serv
}

func (s *Server) LoadEndpoints() {
	s.router.Use(core.LoggingMiddleware)
	s.router.Use(core.ContentTypeMiddleware)
	V1Router := s.router.PathPrefix("/api/v1").Subrouter()
	InternalRouter := s.router.PathPrefix("/internal").Subrouter()
	expressions.LoadEndpoints(V1Router) // api/v1/expressions
	task.LoadEndpoints(InternalRouter)  // internal/task

	s.router.NotFoundHandler = core.WrapHandlerMiddleware(&core.NotFoundHandler{}, core.LoggingMiddleware)
}

func (s *Server) StartServer() {
	s.LoadEndpoints()
	log.Printf("Server started listening on: %s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	log.Fatal(http.ListenAndServe(s.address, s.router))
}
