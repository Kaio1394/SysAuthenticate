package repository

import (
	"SysAuthenticate/internal/models"
	"context"
	"gorm.io/gorm"
)

type UserRoleRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRoleRepositoryImpl(db *gorm.DB) *UserRoleRepositoryImpl {
	return &UserRoleRepositoryImpl{db: db}
}

func (r *UserRoleRepositoryImpl) CreateUserRole(ctx context.Context, userRole *models.UserRole) error {
	return r.db.WithContext(ctx).Create(userRole).Error
}
