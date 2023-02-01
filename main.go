package main

import (
	"portifolio-api/ticket"

	"github.com/gofiber/fiber/v2"
)

func Routers(app *fiber.App) {
	app.Get("/", ticket.GetTickets)
	app.Get("/:id", ticket.GetTicket)
	app.Post("/", ticket.SaveTicket)
	app.Delete("/:id", ticket.DeleteTicket)
	app.Put("/:id", ticket.UpdateTicket)
}

func main() {
	ticket.InitialMigration()
	app := fiber.New()
	Routers(app)
	app.Listen(":3000")
}
