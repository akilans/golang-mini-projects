package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akilans/fiber-book-rest/routes"
	"github.com/gofiber/fiber/v2"
)

// Main function
func main() {
	fmt.Println("Bookstore REST API with MySQL, GORM, JWT, and Fiber")

	// Refer init function defined in models/booksModel.go
	// That loads env, Connects to DB and migrate tables

	// setup app
	app := fiber.New()

	// router config
	routes.Routes(app)

	PORT := os.Getenv("PORT")
	log.Println("Server started on port - ", PORT)
	// start app
	log.Fatal(app.Listen(PORT))
}
