package driven_port

import "github.com/Fei-7/go-hexagonal-api-template/application/model"

type IBookRepository interface {
	GetBooks() []model.Book
}
