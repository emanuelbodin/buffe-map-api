package app

import (
	"os"

	"buffe-map.com/buffe-map-api/config"
	"buffe-map.com/buffe-map-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupAndRunApp() error {
	err := config.LoadEnv()
	if err != nil {
		return err
	}

	app := fiber.New()
	app.Use(logger.New(logger.Config{Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n"}))

	router.SetupRoutes(app)
	config.AddSwaggerRoutes(app)

	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}
