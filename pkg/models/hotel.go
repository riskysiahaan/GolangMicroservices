package models

import (
	"github.com/jinzhu/gorm"
	"github.com/riskysiahaan/proyek_pasti/pkg/config"
)


type Hotel struct {
	gorm.Model
	Nama		string	`json:"nama"`
	Kode 		string 	`json:"kode"`
	Harga		string	`json:"harga"`
	Jenis		string	`json:"jenis"`
	Alamat		string	`json:"alamat"`
	Keterangan	string	`json:"keterangan"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Hotel{})
}

func (b *Hotel) CreateHotel() *Hotel{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllHotels() []Hotel {
	var Hotels []Hotel
	db.Find(&Hotels)
	return Hotels
}

func GetHotelbyId(kode string) (*Hotel, *gorm.DB){
	var getHotel Hotel
	db := db.Where("Kode=?", kode).Find(&getHotel)
	return &getHotel, db
}

func DeleteHotel(kode string) Hotel {
	var hotel Hotel
	db.Where("Kode=?", kode).Delete(hotel)
	return hotel
}