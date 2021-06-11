package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type DefaultRoute struct {
	URI             string
	Method          string
	Function        func(w http.ResponseWriter, r *http.Request)
	IsAuthenticated bool
}

func InjetRoutes(r *mux.Router) *mux.Router {
	routes := UserRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
