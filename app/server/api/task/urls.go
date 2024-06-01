package task

import (
	"github.com/gorilla/mux"
)

func LoadEndpoints(MainRouter *mux.Router) *mux.Router {
	var basePrefix string
	router := MainRouter.PathPrefix("/task").Subrouter()

	expression := &Handler{}
	basePrefix = "task"
	router.HandleFunc("", expression.List).Methods("GET").Name(basePrefix + "-list")
	router.HandleFunc("", expression.Create).Methods("POST").Name(basePrefix + "-create")
	return router
}
