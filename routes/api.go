package routes

import (
	controllers "../controllers"

	"github.com/gorilla/mux"
)

func ApiRoutes(prefix string, r *mux.Router) {
	s := r.PathPrefix(prefix).Subrouter()
	s.Methods("POST").HandlerFunc(controllers.CreateMessage).Path("/messaging")
	s.Methods("GET").HandlerFunc(controllers.SendMessage).Path("/messaging/{id:[0-9]+}")
}
