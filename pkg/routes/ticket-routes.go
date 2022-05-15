package routes

import (
	"github.com/gorilla/mux"
	"github.com/riskysiahaan/proyek_pasti/pkg/controllers"
)

var RegisterTicketsRoutes = func(router *mux.Router) {
	router.HandleFunc("/ticket/",
controllers.CreateTicket).Methods("POST")
	router.HandleFunc("/ticket/", controllers.GetTicket).Methods("GET")
	router.HandleFunc("/ticket/{ticketId}",
controllers.GetTicketById).Methods("GET")
	router.HandleFunc("/ticket/{ticketId}",
controllers.UpdateTicket).Methods("PUT")
	router.HandleFunc("/ticket/{ticketId}",
controllers.DeleteTicket).Methods("DELETE")
}