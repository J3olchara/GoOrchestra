package expressions

import (
	"github.com/gorilla/mux"
)

func LoadEndpoints(MainRouter *mux.Router) *mux.Router {
	var basePrefix string
	router := MainRouter.PathPrefix("/expressions").Subrouter()

	calculate := &Handler{}
	basePrefix = "calculate"
	MainRouter.HandleFunc("/calculate", calculate.Create).Methods("POST").Name(basePrefix + "-create") // mirror for task

	expression := &Handler{}
	basePrefix = "expression"
	router.HandleFunc("", expression.List).Methods("GET").Name(basePrefix + "-list")
	router.HandleFunc("", expression.Create).Methods("POST").Name(basePrefix + "-create")
	router.HandleFunc("/{id}", expression.Retrieve).Methods("GET").Name(basePrefix + "-retrieve")
	return router
}
