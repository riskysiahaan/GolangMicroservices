package routes

import (
	"github.com/gorilla/mux"
	"github.com/riskysiahaan/proyek_pasti/pkg/controllers"
)

var RegisterDestsRoutes = func(router *mux.Router) {
	router.HandleFunc("/destination/",
controllers.CreateDest).Methods("POST")
	router.HandleFunc("/destination/", controllers.GetDest).Methods("GET")
	router.HandleFunc("/destination/{destinationId}",
controllers.GetDestById).Methods("GET")
	router.HandleFunc("/destination/{destinationId}",
controllers.UpdateDest).Methods("PUT")
	router.HandleFunc("/destination/{destinationId}",
controllers.DeleteDest).Methods("DELETE")
}