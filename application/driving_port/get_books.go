package driving_port

import "github.com/Fei-7/go-hexagonal-api-template/application/model"

type IGetBooksHandler interface {
	Handle() []model.Book
}
