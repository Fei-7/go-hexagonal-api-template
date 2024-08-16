package domain

import (
	"github.com/Fei-7/go-hexagonal-api-template/application/driven_port"
	"github.com/Fei-7/go-hexagonal-api-template/application/driving_port"
	"github.com/Fei-7/go-hexagonal-api-template/application/model"
)

type GetBooksHandler struct {
	bookRepository driven_port.IBookRepository
}

func NewGetBooksHandler(bookRepository driven_port.IBookRepository) driving_port.IGetBooksHandler {
	return &GetBooksHandler{
		bookRepository: bookRepository,
	}
}

func (h *GetBooksHandler) Handle() []model.Book {
	books := h.bookRepository.GetBooks()

	return books
}
