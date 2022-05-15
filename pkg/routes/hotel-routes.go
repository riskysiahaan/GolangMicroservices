package routes

import (
	"github.com/gorilla/mux"
	"github.com/riskysiahaan/proyek_pasti/pkg/controllers"
)

var RegisterHotelsRoutes = func(router *mux.Router) {
	router.HandleFunc("/hotel/",
controllers.CreateHotel).Methods("POST")
	router.HandleFunc("/hotel/", controllers.GetPesawat).Methods("GET")
	router.HandleFunc("/hotel/{hotelId}",
controllers.GetHotelById).Methods("GET")
	router.HandleFunc("/hotel/{hotelId}",
controllers.UpdateHotel).Methods("PUT")
	router.HandleFunc("/hotel/{hotelId}",
controllers.DeleteHotel).Methods("DELETE")
}