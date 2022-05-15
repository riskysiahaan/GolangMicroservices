package models

import (
	"github.com/jinzhu/gorm"
	"github.com/riskysiahaan/proyek_pasti/pkg/config"
)

type Dest struct {
	gorm.Model
	Kode		string	`json:"kode"`
	Negara		string	`json:"negara"`
	Provinsi 	string 	`json:"provinsi"`
	Kota		string	`json:"kota"`
	Alamat		string	`json:"alamat"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Dest{})
}

func (b *Dest) CreateDest() *Dest{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllDests() []Dest {
	var Dests []Dest
	db.Find(&Dests)
	return Dests
}

func GetDestbyId(kode int64) (*Dest, *gorm.DB){
	var getDest Dest
	db := db.Where("Kode=?", kode).Find(&getDest)
	return &getDest, db
}

func DeleteDest(kode int64) Dest {
	var destination Dest
	db.Where("Kode=?", kode).Delete(destination)
	return destination
}