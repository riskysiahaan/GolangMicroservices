package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riskysiahaan/proyek_pasti/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTicketsRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}