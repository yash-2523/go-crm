package main

import (
	"fmt"
	"example.com/gofiber-crm/database"
	"example.com/gofiber-crm/lead"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.post("/api/v1/lead", lead.NewLeads)
	app.Delete("/api/v1/lead/:id", lead.DeleteLeads)
}

func initDatabase() {
	var err error
	databse.DBConn, err = gorm.Open("sqlite3", "leads.db")

	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Databse Migrate")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)

	app.Listen(3000)
	defer database.DBConn.Close()
}
