package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/riskysiahaan/proyek_pasti/pkg/models"
	"github.com/riskysiahaan/proyek_pasti/pkg/utils"
)

var NewTicket models.Ticket

func GetTicket(w http.ResponseWriter, r *http.Request) {
	newtickets := models.GetAllTickets()
	res, _ := json.Marshal(newtickets)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTicketById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketId := vars["kode"]
	Kode, err := strconv.ParseInt(ticketId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	ticketDetails, _ := models.GetTicketbyId(Kode)
	res, _ := json.Marshal(ticketDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	CreateTicket := &models.Ticket{}
	utils.ParseBody(r, CreateTicket)
	b := CreateTicket.CreateTicket()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketId := vars["kode"]
	Kode, err := strconv.ParseInt(ticketId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	ticket := models.DeleteTicket(Kode)
	res, _ := json.Marshal(ticket)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	var updateTicket = &models.Ticket{}
	utils.ParseBody(r, updateTicket)
	vars := mux.Vars(r)
	ticketId := vars["kode"]
	Kode, err := strconv.ParseInt(ticketId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	ticketDetails, db := models.GetTicketbyId(Kode)
	if updateTicket.Nama != "" {
	ticketDetails.Nama = updateTicket.Nama
	}
	if updateTicket.Kode != "" {
	ticketDetails.Kode = updateTicket.Kode
	}
	if updateTicket.Harga != "" {
	ticketDetails.Harga = updateTicket.Harga
	}
	if updateTicket.Jenis != "" {
	ticketDetails.Jenis = updateTicket.Jenis
	}
	if updateTicket.Jumlah != "" {
	ticketDetails.Jumlah = updateTicket.Jumlah
	}
	if updateTicket.Keterangan != "" {
	ticketDetails.Keterangan = updateTicket.Keterangan
	}
	db.Save(&ticketDetails)
	res, _ := json.Marshal(ticketDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
	
	
	