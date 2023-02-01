package models

import (
	"gorm.io/gorm"
	"time"
)



type Ticket struct {
	gorm.Model
	Id          uint      `gorm:"primary key;autoIncrement" json:"id"`
	Ticket      int       `json:"ticket"`
	Status      int       `json:"status"`
	Hora_inicio time.Time `json:"hora_inicio"`
	Hora_final  time.Time `json:"hora_final"`
	Dia_semana  int       `json:"dia_semana"`
	Descricao   string    `json:"descricao"`
}

func MigrateTickets(db *gorm.DB) error {
	err := db.AutoMigrate(&Ticket{})
	return err
}
