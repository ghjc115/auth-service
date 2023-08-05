package authService

import (
	"auth-service/dao"
	"auth-service/structs"
	"context"
)

type Service struct {
	userDAO *dao.UserDAO
}

func NewService(userDAO *dao.UserDAO) *Service {
	return &Service{
		userDAO: userDAO,
	}
}

func (s *Service) AddUser(ctx context.Context, req *structs.User) bool {
	return s.userDAO.Add(ctx, req)
}

func (s *Service) GetUser(ctx context.Context, nickname string) *structs.User {
	return s.userDAO.Get(ctx, nickname)
}

func (s *Service) DeleteUser(ctx context.Context, nickname string) bool {
	return s.userDAO.Delete(ctx, nickname)
}
