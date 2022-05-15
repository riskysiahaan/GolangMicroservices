package models

import (
	"github.com/jinzhu/gorm"
	"github.com/riskysiahaan/proyek_pasti/pkg/config"
)

type Pesawat struct {
	gorm.Model
	Nama		string	`json:"nama"`
	Kode 		string 	`json:"kode"`
	Kapasitas	string	`json:"kapasitas"`
	Kelas		string	`json:"kelas"`
	Keterangan	string	`json:"keterangan"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Ticket{})
}

func (b *Pesawat) CreatePesawat() *Pesawat{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllPesawats() []Pesawat {
	var Pesawats []Pesawat
	db.Find(&Pesawats)
	return Pesawats
}

func GetPesawatbyId(kode string) (*Pesawat, *gorm.DB){
	var getPesawat Pesawat
	db := db.Where("Kode=?", kode).Find(&getPesawat)
	return &getPesawat, db
}

func DeletePesawat(kode string) Pesawat {
	var pesawat Pesawat
	db.Where("Kode=?", kode).Delete(pesawat)
	return pesawat
}