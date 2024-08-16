package rest_driving_adapter

import (
	"net/http"

	api_gen "github.com/Fei-7/go-hexagonal-api-template/adapter/driving_adapter/rest_driving_adapter/gen"
	"github.com/gofiber/fiber/v2"
)

func (h *Server) GetBooks(c *fiber.Ctx) error {
	books := h.Domain.GetBooksHandler.Handle()

	responseData := make([]api_gen.Book, 0)
	for _, book := range books {
		responseData = append(responseData, api_gen.Book{
			Id:    &book.Id,
			Title: &book.Title,
		})
	}

	return c.Status(http.StatusOK).JSON(
		api_gen.GetBooksResponse{
			Data: responseData,
		},
	)
}
