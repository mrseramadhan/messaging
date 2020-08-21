package routes

import (
	controllers "../controllers"

	"github.com/gorilla/mux"
)

func ApiRoutes(prefix string, r *mux.Router) {
	s := r.PathPrefix(prefix).Subrouter()
	s.HandleFunc("/messaging/create", controllers.CreateMessage).Methods("POST")
	s.HandleFunc("/messaging/send/{id:[0-9]+}", controllers.SendMessage).Methods("GET")
}
