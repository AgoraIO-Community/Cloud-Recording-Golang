package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/adictya/Agora-Cloud-Recording-Example/api"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	// Set global configuration
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicln(fmt.Errorf("fatal error config file: %s", err))
	}
	viper.AutomaticEnv()

	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", healthCheck)
	api.MountRoutes(app)

	app.Listen(":" + viper.GetString("PORT"))


}
