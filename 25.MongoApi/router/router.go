package router

import (
	"mongoapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllmovies).Methods("GET")

	router.HandleFunc("/api/movies", controller.AddOneMovie).Methods("POST")

	router.HandleFunc("/api/movies/{id}", controller.UpdateWatchedOrNot).Methods("PUT")

	router.HandleFunc("/")

	return router
}
