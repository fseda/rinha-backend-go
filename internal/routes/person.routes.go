package personRoutes

import "github.com/gofiber/fiber/v2"

func SetupPersonRoutes(r *fiber.App) {
	r.Get("/pessoas", )
	r.Get("pessoas/:id", )
	r.Get("/contagem-pessoas")
	r.Post("/pessoas", )
}
