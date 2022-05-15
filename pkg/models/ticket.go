package models

import (
	"github.com/jinzhu/gorm"
	"github.com/riskysiahaan/proyek_pasti/pkg/config"
)

var db *gorm.DB

type Ticket struct {
	gorm.Model
	Nama		string	`json:"nama"`
	Kode 		string 	`json:"kode"`
	Harga		string	`json:"harga"`
	Jenis		string	`json:"jenis"`
	Jumlah		string	`json:"jumlah"`
	Keterangan	string	`json:"keterangan"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Ticket{})
}

func (b *Ticket) CreateTicket() *Ticket{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllTickets() []Ticket {
	var Tickets []Ticket
	db.Find(&Tickets)
	return Tickets
}

func GetTicketbyId(kode int64) (*Ticket, *gorm.DB){
	var getTicket Ticket
	db := db.Where("Kode=?", kode).Find(&getTicket)
	return &getTicket, db
}

func DeleteTicket(kode int64) Ticket {
	var ticket Ticket
	db.Where("Kode=?", kode).Delete(ticket)
	return ticket
}