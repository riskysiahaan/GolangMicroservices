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

var NewDest models.Dest

func GetDest(w http.ResponseWriter, r *http.Request) {
	newdestinations := models.GetAllDests()
	res, _ := json.Marshal(newdestinations)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetDestById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	destinationId := vars["kode"]
	Kode, err := strconv.ParseInt(destinationId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	destinationDetails, _ := models.GetDestbyId(Kode)
	res, _ := json.Marshal(destinationDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateDest(w http.ResponseWriter, r *http.Request) {
	CreateDest := &models.Dest{}
	utils.ParseBody(r, CreateDest)
	b := CreateDest.CreateDest()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteDest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	destinationId := vars["kode"]
	Kode, err := strconv.ParseInt(destinationId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	destination := models.DeleteDest(Kode)
	res, _ := json.Marshal(destination)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateDest(w http.ResponseWriter, r *http.Request) {
	var updateDest = &models.Dest{}
	utils.ParseBody(r, updateDest)
	vars := mux.Vars(r)
	destinationId := vars["kode"]
	Kode, err := strconv.ParseInt(destinationId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	destinationDetails, db := models.GetDestbyId(Kode)
	if updateDest.Negara != "" {
	destinationDetails.Negara = updateDest.Negara
	}
	if updateDest.Kode != "" {
	destinationDetails.Kode = updateDest.Kode
	}
	if updateDest.Provinsi != "" {
	destinationDetails.Provinsi = updateDest.Provinsi
	}
	if updateDest.Kota != "" {
	destinationDetails.Kota = updateDest.Kota
	}
	if updateDest.Alamat != "" {
	destinationDetails.Alamat = updateDest.Alamat
	}
	db.Save(&destinationDetails)
	res, _ := json.Marshal(destinationDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
	
	
	