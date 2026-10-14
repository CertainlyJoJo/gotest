package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// API route
	app.Get("/api/hello", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello from Fiber"})
	})

	app.Use("/", static.New("./public"))

	// Hosts tell you which port to use via $PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
