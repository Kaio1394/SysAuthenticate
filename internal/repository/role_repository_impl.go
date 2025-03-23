package repository

import (
	"SysAuthenticate/internal/models"
	"context"
	"errors"
	"gorm.io/gorm"
)

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepositoryImpl(db *gorm.DB) *RoleRepositoryImpl {
	return &RoleRepositoryImpl{db: db}
}

func (r *RoleRepositoryImpl) CreateNewRole(ctx context.Context, role *models.Role) error {
	return r.db.WithContext(ctx).Create(&role).Error
}

func (r *RoleRepositoryImpl) FindRoleByName(ctx context.Context, nameRole string) (models.Role, error) {
	var role models.Role
	if err := r.db.WithContext(ctx).Where("name = ?", nameRole).First(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

func (r *RoleRepositoryImpl) FindRoleById(ctx context.Context, roleId uint) (*models.Role, error) {
	var role models.Role
	if err := r.db.WithContext(ctx).Where("id = ?", roleId).First(&role).Error; err != nil {
		return nil, errors.New("RoleID not found.")
	}
	return &role, nil
}
