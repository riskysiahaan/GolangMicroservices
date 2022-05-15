package routes

import (
	"github.com/gorilla/mux"
	"github.com/riskysiahaan/proyek_pasti/pkg/controllers"
)

var RegisterPesawatsRoutes = func(router *mux.Router) {
	router.HandleFunc("/pesawat/",
controllers.CreatePesawat).Methods("POST")
	router.HandleFunc("/pesawat/", controllers.GetPesawat).Methods("GET")
	router.HandleFunc("/pesawat/{pesawatId}",
controllers.GetPesawatById).Methods("GET")
	router.HandleFunc("/pesawat/{pesawatId}",
controllers.UpdatePesawat).Methods("PUT")
	router.HandleFunc("/pesawat/{pesawatId}",
controllers.DeletePesawat).Methods("DELETE")
}