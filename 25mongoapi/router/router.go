package router

import (
	"github.com/gorilla/mux"
	"github.com/jay75chauhan/mongoapi/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

    router.HandleFunc("/api/movie",controller.GetMyAllMovie).Methods("GET")
    router.HandleFunc("/api/movi",controller.CreateMovie).Methods("POST")
    router.HandleFunc("/api/movie/{id}",controller.MarkAsWatched).Methods("PUT")
    router.HandleFunc("/api/movie/{id}",controller.DeleteAMovie).Methods("DELETE")
    router.HandleFunc("/api/deleteAllMovie",controller.DeleteALLMovie).Methods("DELETE")

	return router

}