package models

import (
	"time"
	"gorm.io/gorm"
)

// Ticket model
type Ticket struct {
	gorm.Model

	Ticket      int       `json:"ticket"`
	Status      string    `json:"status"`
	Dia_semana  int       `json:"dia_semana"`
	Hora_inicio time.Time `json:"hora_inicio"`
	Hora_final  time.Time `json:"hora_final"`
	Descricao   string    `json:"descricao"`
}
