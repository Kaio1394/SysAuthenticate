package services

import (
	"SysAuthenticate/internal/models"
	"SysAuthenticate/internal/repository"
	"context"
	"errors"
)

type RoleServiceImpl struct {
	r *repository.RoleRepositoryImpl
}

func NewRoleServiceImpl(r *repository.RoleRepositoryImpl) *RoleServiceImpl {
	return &RoleServiceImpl{r: r}
}

func (s *RoleServiceImpl) CreateNewRole(ctx context.Context, role *models.Role) error {
	_, err := s.r.FindRoleByName(ctx, role.Name)
	if err == nil {
		return errors.New("role already exists")
	}

	err = s.r.CreateNewRole(ctx, role)
	if err != nil {
		return err
	}
	return nil
}
