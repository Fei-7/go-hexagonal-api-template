package domain

import (
	"github.com/Fei-7/go-hexagonal-api-template/adapter/driven_adapter/database_adapter"
	"github.com/Fei-7/go-hexagonal-api-template/application/driving_port"
)

type Domains struct {
	GetBooksHandler driving_port.IGetBooksHandler
}

func NewDomains() Domains {
	BookRepository := database_adapter.NewBookRepository()

	GetBooksHandler := NewGetBooksHandler(BookRepository)

	return Domains{
		GetBooksHandler: GetBooksHandler,
	}
}
