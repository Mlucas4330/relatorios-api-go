package routes

import (
	"portifolio-api/database"
	"portifolio-api/models"
	"github.com/gofiber/fiber/v2"
)

// Book
func Book(c *fiber.Ctx) error {
	tickets := []models.Ticket{}
	ticket := new(models.Ticket)
	if err := c.BodyParser(ticket); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Where("title = ?", ticket.Ticket).Find(&tickets)
	return c.Status(200).JSON(tickets)
}

// AddBook
func AddBook(c *fiber.Ctx) error {
	ticket := new(models.Ticket)
	if err := c.BodyParser(ticket); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Create(&ticket)

	return c.Status(200).JSON(ticket)
}

// AllBooks
func AllBooks(c *fiber.Ctx) error {
	tickets := []models.Ticket{}
	database.DB.Db.Find(&tickets)

	return c.Status(200).JSON(tickets)
}

// Update
func Update(c *fiber.Ctx) error {
	tickets := []models.Ticket{}
	ticket := new(models.Ticket)
	if err := c.BodyParser(ticket); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Db.Model(&tickets).Where("title = ?", ticket.Ticket).Update("author", ticket)

	return c.Status(400).JSON("updated")
}

// Delete
func Delete(c *fiber.Ctx) error {
	book := []models.Book{}
	title := new(models.Book)
	if err := c.BodyParser(title); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Db.Where("title = ?", title.Title).Delete(&book)

	return c.Status(200).JSON("deleted")
}
