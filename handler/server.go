package handler

import (
	"github.com/SawitProRecruitment/UserService/jwt_usecase"
	"github.com/SawitProRecruitment/UserService/repository"
)

type Server struct {
	Repository repository.RepositoryInterface
	JWTUsecase jwt_usecase.JWTUsecaseInterface
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
	JWTUsecase jwt_usecase.JWTUsecaseInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Repository: opts.Repository,
		JWTUsecase: opts.JWTUsecase,
	}
}
