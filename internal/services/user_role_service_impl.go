package services

import (
	"SysAuthenticate/internal/models"
	"SysAuthenticate/internal/repository"
	"context"
)

type UserRoleServiceImpl struct {
	r  *repository.UserRoleRepositoryImpl
	rr *repository.RoleRepositoryImpl
	rs *repository.UserRepositoryImpl
}

func NewUserRoleServiceImpl(r *repository.UserRoleRepositoryImpl) *UserRoleServiceImpl {
	return &UserRoleServiceImpl{r: r}
}

func (s *UserRoleServiceImpl) CreateUserRole(ctx context.Context, userRole *models.UserRole) error {
	return s.r.CreateUserRole(ctx, userRole)
}
