package app

import (
	"os"

	"github.com/fseda/rinha-backend-go/config"
	"github.com/fseda/rinha-backend-go/database"
	"github.com/fseda/rinha-backend-go/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupAndRunApp() error {
	var err error

	err = config.LoadENV()
	if err != nil {
		return err
	}

	err = database.InitializeDB(os.Getenv("DB_CONN_STR"))
	if err != nil {
		return err
	}
	defer database.CloseDB()

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] [${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	router.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err = app.Listen(":" + port)
	return nil
}
