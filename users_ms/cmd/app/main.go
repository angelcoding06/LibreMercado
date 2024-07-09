package main

import (
	"github.com/angelcoding06/LibreMercado/users_ms/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")

	})

	log.Printf("Server running on port %s", cfg.Port)
	app.Listen(":" + cfg.Port)

}
