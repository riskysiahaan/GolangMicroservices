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

var NewPesawat models.Pesawat

func GetPesawat(w http.ResponseWriter, r *http.Request) {
	newpesawats := models.GetAllPesawats()
	res, _ := json.Marshal(newpesawats)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPesawatById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pesawatId := vars["pesawatId"]
	Kode, err := strconv.ParseInt(pesawatId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	pesawatDetails, _ := models.GetPesawatbyId(Kode)
	res, _ := json.Marshal(pesawatDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatePesawat(w http.ResponseWriter, r *http.Request) {
	CreatePesawat := &models.Pesawat{}
	utils.ParseBody(r, CreatePesawat)
	b := CreatePesawat.CreatePesawat()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePesawat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pesawatId := vars["pesawatId"]
	Kode, err := strconv.ParseInt(pesawatId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	pesawat := models.DeleteTicket(Kode)
	res, _ := json.Marshal(pesawat)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePesawat(w http.ResponseWriter, r *http.Request) {
	var updatePesawat = &models.Pesawat{}
	utils.ParseBody(r, updatePesawat)
	vars := mux.Vars(r)
	pesawatId := vars["pesawatId"]
	Kode, err := strconv.ParseInt(pesawatId, 0, 0)
	if err != nil {
	fmt.Println("error while parsing")
	}
	pesawatDetails, db := models.GetPesawatbyId(Kode)
	if updatePesawat.Nama != "" {
	pesawatDetails.Nama = updatePesawat.Nama
	}
	if updatePesawat.Kode != "" {
	pesawatDetails.Kode = updatePesawat.Kode
	}
	if updatePesawat.Kapasitas != "" {
	pesawatDetails.Kapasitas = updatePesawat.Kapasitas
	}
	if updatePesawat.Kelas != "" {
	pesawatDetails.Kelas = updatePesawat.Kelas
	}
	if updatePesawat.Keterangan != "" {
	pesawatDetails.Keterangan = updatePesawat.Keterangan
	}
	db.Save(&pesawatDetails)
	res, _ := json.Marshal(pesawatDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
	
	
	