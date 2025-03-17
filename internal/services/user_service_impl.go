package services

import (
	"SysAuthenticate/internal/models/dtos/read"
	"SysAuthenticate/internal/repository"
	"context"
	"github.com/jinzhu/copier"
)

type UserServiceImpl struct {
	r *repository.UserRepositoryImpl
}

func NewUserServiceImpl(r *repository.UserRepositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{r: r}
}

func (s *UserServiceImpl) GetUsers(ctx context.Context) ([]read.UserReadDto, error) {
	var usersDto []read.UserReadDto
	listUsers, err := s.r.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	_ = copier.CopyWithOption(&usersDto, &listUsers, copier.Option{IgnoreEmpty: true})
	return usersDto, nil
}
