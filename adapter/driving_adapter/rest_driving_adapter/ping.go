package rest_driving_adapter

import (
	"net/http"

	api_gen "github.com/Fei-7/go-hexagonal-api-template/adapter/driving_adapter/rest_driving_adapter/gen"
	"github.com/gofiber/fiber/v2"
)

func (h *Server) GetPing(ctx *fiber.Ctx) error {
	resp := api_gen.Pong{
		Ping: "pong",
	}

	return ctx.
		Status(http.StatusOK).
		JSON(resp)
}
