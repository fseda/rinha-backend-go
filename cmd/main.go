package main

import (
	"github.com/fseda/rinha-backend-go/config"
	"github.com/fseda/rinha-backend-go/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "Rinha-Backend-Go",
	})

	personRoutes.SetupPersonRoutes(app)

	port, _ := config.Config("PORT")
	app.Listen(":"+port)
}
