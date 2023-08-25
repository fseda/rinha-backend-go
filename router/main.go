package router

import (
	"github.com/fseda/rinha-backend-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/pessoas", handlers.HandleCreatePerson)
	app.Get("/pessoas", handlers.HandleGetPersonById)
	app.Get("/pessoas/:id", handlers.HandleGetPersonById)
	app.Get("contagem-pessoas", handlers.HandleCountPeople)
}