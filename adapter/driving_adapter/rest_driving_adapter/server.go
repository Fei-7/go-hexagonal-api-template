package rest_driving_adapter

import (
	api_gen "github.com/Fei-7/go-hexagonal-api-template/adapter/driving_adapter/rest_driving_adapter/gen"
	"github.com/Fei-7/go-hexagonal-api-template/application/domain"
)

var _ api_gen.ServerInterface = (*Server)(nil)

type Server struct {
	Domain domain.Domains
}

func NewServer() *Server {
	return &Server{
		Domain: domain.NewDomains(),
	}
}
