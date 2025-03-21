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
	_, err := s.rs.GetUserById(ctx, userRole.UserID)
	if err != nil {
		return err
	}
	_, err = s.rr.FindRoleById(ctx, userRole.RoleID)
	if err != nil {
		return err
	}
	return s.r.CreateUserRole(ctx, userRole)
}
