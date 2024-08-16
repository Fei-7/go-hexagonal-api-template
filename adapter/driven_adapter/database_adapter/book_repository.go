package database_adapter

import (
	"github.com/Fei-7/go-hexagonal-api-template/application/driven_port"
	"github.com/Fei-7/go-hexagonal-api-template/application/model"
)

type BookRepository struct {
}

func NewBookRepository() driven_port.IBookRepository {
	return &BookRepository{}
}

func (h *BookRepository) GetBooks() []model.Book {
	return []model.Book{
		{
			Id:    "1",
			Title: "Mock Book Data",
		},
	}
}
