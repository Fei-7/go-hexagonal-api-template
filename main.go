package main

import (
	"log"

	"github.com/Fei-7/go-hexagonal-api-template/adapter/driving_adapter/rest_driving_adapter"
	api_gen "github.com/Fei-7/go-hexagonal-api-template/adapter/driving_adapter/rest_driving_adapter/gen"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	server := rest_driving_adapter.NewServer()

	app := fiber.New()

	app.Use(logger.New())

	api_gen.RegisterHandlers(app, server)

	log.Fatal(app.Listen("0.0.0.0:8080"))
}
