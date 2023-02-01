package ticket

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:admin@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

type Ticket struct {
	gorm.Model
	Ticket      *int       `json:"ticket"`
	Status      int       `json:"status"`
	Hora_inicio time.Time `json:"hora_inicio"`
	Hora_final  time.Time `json:"hora_final"`
	Dia_semana  int       `json:"dia_semana"`
	Descricao   string    `json:"descricao"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&Ticket{})
}

func GetTickets(c *fiber.Ctx) error {
	var tickets []Ticket
	DB.Find(&tickets)
	return c.JSON(&tickets)
}

func GetTicket(c *fiber.Ctx) error {
	id := c.Params("id")
	var ticket Ticket
	DB.Find(&ticket, id)
	return c.JSON(&ticket)
}

func SaveTicket(c *fiber.Ctx) error {
	ticket := new(Ticket)
	if err := c.BodyParser(ticket); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&ticket)
	return c.JSON(&ticket)
}

func DeleteTicket(c *fiber.Ctx) error {
	id := c.Params("id")
	var ticket Ticket
	DB.First(&ticket, id)
	if ticket.Ticket == nil {
		return c.Status(500).SendString("Ticket not available")
	}

	DB.Delete(&ticket)
	return c.SendString("Ticket is deleted!")
}

func UpdateTicket(c *fiber.Ctx) error {
	id := c.Params("id")
	ticket := new(Ticket)
	DB.First(&ticket, id)
	if ticket.Ticket == nil {
		return c.Status(500).SendString("Ticket not available")
	}
	if err := c.BodyParser(&ticket); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&ticket)
	return c.JSON(&ticket)
}
