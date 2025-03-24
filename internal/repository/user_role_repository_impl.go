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

func (r *UserRoleRepositoryImpl) GetUserRoleById(ctx context.Context, id uint) (models.UserRole, error) {
	var userRole models.UserRole
	if err := r.db.WithContext(ctx).Where("user_id=?", id).First(&userRole).Error; err != nil {
		return models.UserRole{}, err
	}
	return userRole, nil
}
