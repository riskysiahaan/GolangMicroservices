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

var NewHotel models.Hotel

func GetHotel(w http.ResponseWriter, r *http.Request) {
	newhotels := models.GetAllHotels()
	res, _ := json.Marshal(newhotels)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetHotelById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hotelId := vars["kode"]
	Kode, err := strconv.ParseInt(hotelId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	hotelDetails, _ := models.GetHotelbyId(Kode)
	res, _ := json.Marshal(hotelDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	CreateHotel := &models.Hotel{}
	utils.ParseBody(r, CreateHotel)
	b := CreateHotel.CreateHotel()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hotelId := vars["kode"]
	Kode, err := strconv.ParseInt(hotelId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	hotel := models.DeleteHotel(Kode)
	res, _ := json.Marshal(hotel)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	var updateHotel = &models.Hotel{}
	utils.ParseBody(r, updateHotel)
	vars := mux.Vars(r)
	hotelId := vars["kode"]
	Kode, err := strconv.ParseInt(hotelId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	hotelDetails, db := models.GetHotelbyId(Kode)
	if updateHotel.Nama != "" {
	hotelDetails.Nama = updateHotel.Nama
	}
	if updateHotel.Kode != "" {
	hotelDetails.Kode = updateHotel.Kode
	}
	if updateHotel.Harga != "" {
	hotelDetails.Harga = updateHotel.Harga
	}
	if updateHotel.Jenis != "" {
	hotelDetails.Jenis = updateHotel.Jenis
	}
	if updateHotel.Alamat != "" {
	hotelDetails.Alamat = updateHotel.Alamat
	}
	if updateHotel.Keterangan != "" {
	hotelDetails.Keterangan = updateHotel.Keterangan
	}
	db.Save(&hotelDetails)
	res, _ := json.Marshal(hotelDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
	
	
	